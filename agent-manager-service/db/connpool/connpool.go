// Copyright (c) 2025, WSO2 LLC (http://www.wso2.com). All Rights Reserved.
//
// This software is the property of WSO2 LLC and its suppliers, if any.
// Dissemination of any information or reproduction of any material contained
// herein is strictly forbidden, unless permitted by WSO2 in accordance with
// the WSO2 Commercial License available at http://wso2.com/licenses.
// For specific language governing the permissions and limitations under
// this license, please see the license as well as any agreement you've
// entered into with WSO2 governing the purchase of this software and any
// associated services.

package connpool

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"
	"time"

	"gorm.io/gorm"
)

var WaitFunc = func(d time.Duration) {
	<-time.After(d)
}

type BackoffFunc func(failedCount int) time.Duration

type RetryParams struct {
	MaxRetries  int
	BackoffFunc BackoffFunc
}

// a wrapper to add a retry mechanism to sql.DB methods
type connPool struct {
	*sql.DB
	retryParams RetryParams
}

var (
	_ gorm.ConnPool       = &connPool{}
	_ gorm.GetDBConnector = &connPool{}
)

func New(db *sql.DB, retryParams RetryParams) gorm.ConnPool {
	if retryParams.BackoffFunc == nil {
		slog.Info("setting default backoff function", "log_type", "connPool")
		retryParams.BackoffFunc = func(failedCount int) time.Duration {
			return time.Millisecond * 200
		}
	}
	return &connPool{db, retryParams}
}

func (c *connPool) GetDBConn() (*sql.DB, error) {
	return c.DB, nil
}

func (c *connPool) Conn(ctx context.Context) (*sql.Conn, error) {
	slog.Debug("connPool:Conn", "log_type", "connPool")
	return c.DB.Conn(ctx)
}

func (c *connPool) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sql.Tx, err error) {
	slog.Debug("connPool:BeginTx", "log_type", "connPool", "opts", opts)
	c.retry(ctx, "BeginTx", func() error {
		// `tx` contains *sql.DB, retries won't work for queries run on `tx`
		tx, err = c.DB.BeginTx(ctx, opts)
		return err
	})
	return tx, err
}

func (c *connPool) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	slog.Debug("connPool:PrepareContext", "log_type", "connPool", "query", query)
	c.retry(ctx, "PrepareContext", func() error {
		// `stmt` contains *sql.DB, retries won't work for queries run on `stmt`
		stmt, err = c.DB.PrepareContext(ctx, query)
		return err
	})
	return stmt, err
}

func (c *connPool) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	slog.Debug("connPool:ExecContext", "log_type", "connPool", "query", query)
	c.retry(ctx, "ExecContext", func() error {
		result, err = c.DB.ExecContext(ctx, query, args...)
		return err
	})
	return result, err
}

func (c *connPool) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	slog.Debug("connPool:QueryContext", "log_type", "connPool", "query", query)
	c.retry(ctx, "QueryContext", func() error {
		rows, err = c.DB.QueryContext(ctx, query, args...)
		return err
	})
	return rows, err
}

func (c *connPool) QueryRowContext(ctx context.Context, query string, args ...interface{}) (val *sql.Row) {
	slog.Debug("connPool:QueryRowContext", "log_type", "connPool", "query", query)
	c.retry(ctx, "QueryRowContext", func() error {
		val = c.DB.QueryRowContext(ctx, query, args...)
		return val.Err()
	})
	return val
}

func (c *connPool) retry(ctx context.Context, fn string, op func() error) {
	var err error
	for attempts := 1; attempts <= c.retryParams.MaxRetries; attempts++ {
		// Check context before attempting operation
		select {
		// Context was cancelled (timeout/user cancellation)
		case <-ctx.Done():
			slog.Warn("connPool operation canceled",
				"log_type", "connPool",
				"fn", fn,
				"context_error", ctx.Err())
			return
		default:
		}

		err = op()
		if err == nil || !isRetryableError(err) {
			return
		}

		slog.Error("connPool operation failed",
			"log_type", "connPool",
			"error", err,
			"attempt", attempts,
			"fn", fn)

		if attempts < c.retryParams.MaxRetries {
			backoffDuration := c.retryParams.BackoffFunc(attempts)

			select {
			case <-time.After(backoffDuration):
				// Wait completed, continue to next retry
			case <-ctx.Done():
				slog.Warn("connPool operation canceled during backoff",
					"log_type", "connPool",
					"fn", fn,
					"context_error", ctx.Err())
				return
			}
		}
	}
}

func isRetryableError(err error) bool {
	msg := err.Error()
	// read tcp 172.17.1.242:33978->172.17.64.21:1433: i/o timeout
	if strings.Contains(msg, "read tcp") && strings.Contains(msg, "i/o timeout") {
		return true
	}
	for _, s := range []string{dbIsTimeout, dbIsClosed, dbIsFailedRPC} {
		if strings.Contains(msg, s) {
			return true
		}
	}
	return false
}

// https://github.com/wso2-enterprise/choreo-connect-global-adapter/pull/201/files
const (
	dbIsClosed    = "is closed"
	dbIsTimeout   = "timed out"
	dbIsFailedRPC = "failed to send RPC"
)
