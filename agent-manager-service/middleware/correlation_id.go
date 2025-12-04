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
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/utils"
)

const (
	CorrelationIDHeader string = "x-correlation-id"
)

// AddCorrelationID middleware adds or generates a correlation ID for request tracing
func AddCorrelationID() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get or generate correlation ID
			correlationID := r.Header.Get(CorrelationIDHeader)
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			// Set response header
			w.Header().Set(CorrelationIDHeader, correlationID)

			// Add to request context
			ctx := context.WithValue(r.Context(), utils.CorrelationIdCtxKey(), correlationID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
