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

package jwtassertion

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

// NewMockMiddleware creates a mock JWT middleware for testing
func NewMockMiddleware(t *testing.T, orgId uuid.UUID, userIdpId uuid.UUID) Middleware {
	t.Helper()

	tokenClaims := &TokenClaims{
		Sub:   userIdpId,
		Scope: "scopes",
		Exp:   int(time.Now().Add(time.Hour).Unix()),
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Set the context values that GetTokenClaims expects
			ctx = context.WithValue(ctx, assertionTokenClaimsKey, tokenClaims)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
