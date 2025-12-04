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
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
)

type ctxTX struct{}

func DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(ctxTX{}).(*gorm.DB)
	if ok {
		return tx
	}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		timeoutCtx, cancel := context.WithTimeout(ctx,
			time.Duration(config.GetConfig().DbOperationTimeoutSeconds)*time.Second)
		// Note: We don't defer cancel() here because the returned *gorm.DB
		// will be used beyond this function's scope.
		_ = cancel
		return db.WithContext(timeoutCtx)
	}

	return db.WithContext(ctx)
}

func CtxWithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxTX{}, tx)
}

func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
