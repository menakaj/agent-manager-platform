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

package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/utils"
)

func RecovererOnPanic() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					correlationId := "unknown"
					if id := r.Context().Value(utils.CorrelationIdCtxKey()); id != nil {
						if idStr, ok := id.(string); ok {
							correlationId = idStr
						}
					}

					operation := "unknown"
					if op := r.Context().Value("operation"); op != nil {
						if opStr, ok := op.(string); ok {
							operation = opStr
						}
					}

					slog.Error("recoverOnPanic",
						"correlationID", correlationId,
						"operation", operation,
						"log_type", "err_response",
						"panic", rec,
						"stack", string(debug.Stack()))

					utils.WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
