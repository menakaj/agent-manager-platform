//go:build wireinject
// +build wireinject

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
	"log/slog"

	"github.com/google/wire"

	observabilitysvc "github.com/wso2-enterprise/agent-management-platform/agent-manager-service/clients/observabilitysvc"
	clients "github.com/wso2-enterprise/agent-management-platform/agent-manager-service/clients/openchoreosvc"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/controllers"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/jwtassertion"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/repositories"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/services"
)

var configProviderSet = wire.NewSet(
	ProvideConfigFromPtr,
)

var repositoryProviderSet = wire.NewSet(
	repositories.NewOrganizationRepository,
	repositories.NewAgentRepository,
	repositories.NewProjectRepository,
)

var clientProviderSet = wire.NewSet(
	clients.NewOpenChoreoSvcClient,
	observabilitysvc.NewObservabilitySvcClient,
)

var serviceProviderSet = wire.NewSet(
	services.NewAgentManagerService,
	services.NewBuildCIManager,
	services.NewInfraResourceManager,
)

var controllerProviderSet = wire.NewSet(
	controllers.NewAgentController,
	controllers.NewBuildCIController,
	controllers.NewInfraResourceController,
)

var testClientProviderSet = wire.NewSet(
	ProvideTestOpenChoreoSvcClient,
	ProvideTestObservabilitySvcClient,
)

// ProvideLogger provides the configured slog.Logger instance
func ProvideLogger() *slog.Logger {
	return slog.Default()
}

var loggerProviderSet = wire.NewSet(
	ProvideLogger,
)

// ProvideTestOpenChoreoSvcClient extracts the OpenChoreoSvcClient from TestClients
func ProvideTestOpenChoreoSvcClient(testClients TestClients) clients.OpenChoreoSvcClient {
	return testClients.OpenChoreoSvcClient
}

// ProvideTestObservabilitySvcClient extracts the ObservabilitySvcClient from TestClients
func ProvideTestObservabilitySvcClient(testClients TestClients) observabilitysvc.ObservabilitySvcClient {
	return testClients.ObservabilitySvcClient
}

func InitializeAppParams(cfg *config.Config) (*AppParams, error) {
	wire.Build(
		configProviderSet,
		repositoryProviderSet,
		clientProviderSet,
		loggerProviderSet,
		serviceProviderSet,
		controllerProviderSet,
		ProvideAuthMiddleware,
		wire.Struct(new(AppParams), "*"),
	)
	return &AppParams{}, nil
}

func InitializeTestAppParamsWithClientMocks(cfg *config.Config, authMiddleware jwtassertion.Middleware, testClients TestClients) (*AppParams, error) {
	wire.Build(
		repositoryProviderSet,
		testClientProviderSet,
		loggerProviderSet,
		serviceProviderSet,
		controllerProviderSet,
		wire.Struct(new(AppParams), "*"),
	)
	return &AppParams{}, nil
}
