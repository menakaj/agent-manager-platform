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

package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/db/connpool"
)

var db *gorm.DB

func init() {
	db = initDbConn(config.GetConfig().POSTGRESQL)
}

// slogWriter implements the GORM logger Writer interface using slog
type slogWriter struct{}

func (w slogWriter) Printf(format string, v ...interface{}) {
	slog.Warn(fmt.Sprintf(format, v...), "log_type", "gorm")
}

func initDbConn(cfg config.POSTGRESQL) *gorm.DB {
	dsn := makeConnString(cfg)
	sqlConnPool, err := sql.Open("pgx", dsn)
	if err != nil {
		slog.Error("initDbConn: sql.Open failed", "error", err)
		os.Exit(1)
	}
	setConfigsOnDB(sqlConnPool, cfg.DbConfigs)
	if err := sqlConnPool.Ping(); err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(1)
	}
	connPool := connpool.New(sqlConnPool, connpool.RetryParams{
		MaxRetries: 3,
		BackoffFunc: func(failedCount int) time.Duration {
			return time.Duration(failedCount) * time.Second * 5
		},
	})

	gormLogger := logger.New(
		slogWriter{},
		logger.Config{
			SlowThreshold:             time.Duration(cfg.SlowThresholdMilliseconds) * time.Millisecond,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		},
	)

	// Open PostgreSQL connection
	dialector := postgres.Open(dsn)
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: cfg.SkipDefaultTransaction,
		PrepareStmt:            false,
		FullSaveAssociations:   false,
		ConnPool:               connPool,
	})
	if err != nil {
		slog.Error("initDbConn: gorm.Open failed", "error", err)
		os.Exit(1)
	}
	slog.Info("database connected")
	return gormDB
}

func setConfigsOnDB(db *sql.DB, cfg config.DbConfigs) {
	if cfg.MaxIdleTimeSeconds != nil {
		db.SetConnMaxIdleTime(time.Duration(*cfg.MaxIdleTimeSeconds) * time.Second)
	}
	if cfg.MaxLifetimeSeconds != nil {
		db.SetConnMaxLifetime(time.Duration(*cfg.MaxLifetimeSeconds) * time.Second)
	}
	if cfg.MaxOpenCount != nil {
		db.SetMaxOpenConns(int(*cfg.MaxOpenCount))
	}
	if cfg.MaxIdleCount != nil {
		db.SetMaxIdleConns(int(*cfg.MaxIdleCount))
	}
}

func makeConnString(p config.POSTGRESQL) string {
	params := url.Values{}
	conn := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(p.User, p.Password),
		Host:     fmt.Sprintf("%s:%d", p.Host, p.Port),
		Path:     "/" + p.DBName,
		RawQuery: params.Encode(),
	}
	return conn.String()
}
