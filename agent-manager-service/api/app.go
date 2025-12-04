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

package api

import (
	"net/http"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/logger"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/wiring"
)

// MakeHTTPHandler creates a new HTTP handler with middleware and routes
func MakeHTTPHandler(params *wiring.AppParams) http.Handler {
	mux := http.NewServeMux()

	// Register health check
	registerHealthCheck(mux)

	// Create a sub-mux for API v1 routes
	apiMux := http.NewServeMux()
	registerAgentRoutes(apiMux, params.AgentController)
	registerInfraRoutes(apiMux, params.InfraResourceController)

	// Apply middleware in reverse order (last middleware is applied first)
	apiHandler := http.Handler(apiMux)
	apiHandler = middleware.RecovererOnPanic()(apiHandler)
	apiHandler = logger.RequestLogger()(apiHandler)
	apiHandler = params.AuthMiddleware(apiHandler)
	apiHandler = middleware.AddCorrelationID()(apiHandler)

	// Create a mux for internal API routes
	internalApiMux := http.NewServeMux()
	registerInternalRoutes(internalApiMux, params.BuildCIController)
	internalApiHandler := http.Handler(internalApiMux)
	internalApiHandler = middleware.RecovererOnPanic()(internalApiHandler)
	internalApiHandler = logger.RequestLogger()(internalApiHandler)
	internalApiHandler = middleware.APIKeyMiddleware()(internalApiHandler) // Add API key middleware for internal routes temporarily
	internalApiHandler = middleware.AddCorrelationID()(internalApiHandler)
	apiHandler = middleware.CORS(config.GetConfig().CORSAllowedOrigin)(apiHandler)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiHandler))
	mux.Handle("/internal/", http.StripPrefix("/internal", internalApiHandler))

	return mux
}
