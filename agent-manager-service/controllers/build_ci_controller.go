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

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/logger"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/services"
)

type BuildCallbackPayload struct {
	ImageID     string `json:"imageId"`
	AgentName   string `json:"agentName"`
	ProjectName string `json:"projectName"`
	OrgName     string `json:"orgName"`
}

type BuildCIController interface {
	HandleBuildCallback(w http.ResponseWriter, r *http.Request)
}

type buildCIController struct {
	buildCIManagerService services.BuildCIManagerService
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func NewBuildCIController(buildCIManagerService services.BuildCIManagerService) BuildCIController {
	return &buildCIController{
		buildCIManagerService: buildCIManagerService,
	}
}

func (b *buildCIController) HandleBuildCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	var payload BuildCallbackPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error("HandleBuildCallback: failed to decode request body", "error", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	log.Info("Build callback received", "payload", payload)

	b.buildCIManagerService.HandleBuildCallback(ctx, payload.AgentName, payload.ProjectName, payload.OrgName, payload.ImageID)
	writeJSONResponse(w, http.StatusAccepted, map[string]string{"status": "accepted"})
}
