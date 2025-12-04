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

package logger

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/utils"
)

type loggerKey struct{}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger retrieves the logger from context, or returns the configured default logger
func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}
	// Use the globally configured logger instead of slog.Default()
	return slog.Default()
}

func RequestLogger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			correlationID := "unknown"
			if id := r.Context().Value(utils.CorrelationIdCtxKey()); id != nil {
				if idStr, ok := id.(string); ok {
					correlationID = idStr
				}
			}
			// Use the globally configured logger
			reqLogger := slog.Default().With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("correlation_id", correlationID),
			)
			ctx := WithLogger(r.Context(), reqLogger)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
