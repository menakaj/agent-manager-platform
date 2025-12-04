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

func registerInfraRoutes(mux *http.ServeMux, ctrl controllers.InfraResourceController) {
	mux.HandleFunc("POST /orgs", ctrl.CreateOrganization)
	mux.HandleFunc("GET /orgs", ctrl.ListOrganizations)
	mux.HandleFunc("GET /orgs/{orgName}", ctrl.GetOrganization)
	mux.HandleFunc("GET /orgs/{orgName}/projects", ctrl.ListProjects)
	mux.HandleFunc("POST /orgs/{orgName}/projects", ctrl.CreateProject)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projectName}", ctrl.GetProject)
	mux.HandleFunc("GET /orgs/{orgName}/environments", ctrl.GetOrgEnvironments)
	mux.HandleFunc("GET /orgs/{orgName}/projects/{projectName}/deployment-pipelines", ctrl.GetProjectDeploymentPipeline)
}
