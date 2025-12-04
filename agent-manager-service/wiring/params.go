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

package wiring

import (
	observabilitysvc "github.com/wso2-enterprise/agent-management-platform/agent-manager-service/clients/observabilitysvc"
	clients "github.com/wso2-enterprise/agent-management-platform/agent-manager-service/clients/openchoreosvc"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/controllers"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/jwtassertion"
)

type AppParams struct {
	AuthMiddleware          jwtassertion.Middleware
	AgentController         controllers.AgentController
	InfraResourceController controllers.InfraResourceController
	BuildCIController       controllers.BuildCIController
}

// TestClients contains all mock clients needed for testing
type TestClients struct {
	OpenChoreoSvcClient    clients.OpenChoreoSvcClient
	ObservabilitySvcClient observabilitysvc.ObservabilitySvcClient
}

func ProvideConfigFromPtr(config *config.Config) config.Config {
	return *config
}

func ProvideAuthMiddleware(config config.Config) jwtassertion.Middleware {
	return jwtassertion.JWTAuthMiddleware(config.AuthHeader)
}
