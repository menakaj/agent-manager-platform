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

package services

import (
	"context"
	"log/slog"

	clients "github.com/wso2-enterprise/agent-management-platform/agent-manager-service/clients/openchoreosvc"
)

type BuildCIManagerService interface {
	HandleBuildCallback(ctx context.Context, agentName string, projectName string, orgName string, imageId string)
}

type buildCIManagerService struct {
	OpenChoreoSvcClient clients.OpenChoreoSvcClient
	logger              *slog.Logger
}

func NewBuildCIManager(openChoreoSvcClient clients.OpenChoreoSvcClient, logger *slog.Logger,
) BuildCIManagerService {
	return &buildCIManagerService{
		OpenChoreoSvcClient: openChoreoSvcClient,
		logger:              logger,
	}
}

func (b *buildCIManagerService) HandleBuildCallback(ctx context.Context, agentName string, projectName string, orgName string, imageId string) {
	_, err := b.OpenChoreoSvcClient.GetProject(ctx, projectName, orgName)
	if err != nil {
		b.logger.Error("Project not found", "project", projectName, "organization", orgName)
		return
	}

	component, err := b.OpenChoreoSvcClient.GetAgentComponent(ctx, orgName, projectName, agentName)
	if err != nil {
		b.logger.Error("Failed to get component", "component", agentName, "project", projectName, "organization", orgName, "error", err)
		return
	}
	if err := b.OpenChoreoSvcClient.DeployBuiltImage(ctx, orgName, projectName, component.Name, imageId); err != nil {
		b.logger.Error("Failed to deploy agent component", "component", component.Name, "error", err)
		return
	}
}
