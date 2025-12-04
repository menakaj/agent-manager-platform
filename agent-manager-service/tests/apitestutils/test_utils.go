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

package apitestutils

import (
	"net/http"
	"testing"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/api"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/jwtassertion"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/wiring"
)

// MakeAppClientWithDeps creates an HTTP handler with the provided dependencies for testing
func MakeAppClientWithDeps(t *testing.T, testClients wiring.TestClients, authMiddleware jwtassertion.Middleware) http.Handler {
	// Use wire to initialize the app parameters with test clients
	appParams, err := wiring.InitializeTestAppParamsWithClientMocks(config.GetConfig(), authMiddleware, testClients)
	if err != nil {
		t.Fatalf("failed to initialize test app params: %v", err)
	}

	// Create HTTP handler
	handler := api.MakeHTTPHandler(appParams)

	// Return the handler instance
	return handler
}
