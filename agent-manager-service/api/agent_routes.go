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

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/controllers"
)

func registerAgentRoutes(mux *http.ServeMux, ctrl controllers.AgentController) {
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents", ctrl.ListAgents)
	// utility endpoint to generate name from  name
	mux.HandleFunc("POST /orgs/{orgName}/utils/generate-name", ctrl.GenerateName)
	mux.HandleFunc("POST /orgs/{orgName}/projects/{projName}/agents", ctrl.CreateAgent)

	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}", ctrl.GetAgent)
	mux.HandleFunc("DELETE /orgs/{orgName}/projects/{projName}/agents/{agentName}", ctrl.DeleteAgent)
	mux.HandleFunc("POST /orgs/{orgName}/projects/{projName}/agents/{agentName}/builds", ctrl.BuildAgent)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/builds", ctrl.ListAgentBuilds)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/builds/{buildName}", ctrl.GetBuild)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/builds/{buildName}/build-logs", ctrl.GetBuildLogs)

	mux.HandleFunc("POST /orgs/{orgName}/projects/{projName}/agents/{agentName}/deployments", ctrl.DeployAgent)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/deployments", ctrl.GetAgentDeployments)

	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/endpoints", ctrl.GetAgentEndpoints)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projName}/agents/{agentName}/configurations", ctrl.GetAgentConfigurations)
}
