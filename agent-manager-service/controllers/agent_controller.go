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

// NOTE: This is a placeholder implementation for the agent controller.
// TODO: Replace with actual implementation based on your HTTP framework (Gin, Echo, etc.).

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/jwtassertion"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/middleware/logger"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/services"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/spec"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/utils"
)

type AgentController interface {
	ListAgents(w http.ResponseWriter, r *http.Request)
	GetAgent(w http.ResponseWriter, r *http.Request)
	CreateAgent(w http.ResponseWriter, r *http.Request)
	DeleteAgent(w http.ResponseWriter, r *http.Request)
	BuildAgent(w http.ResponseWriter, r *http.Request)
	DeployAgent(w http.ResponseWriter, r *http.Request)
	ListAgentBuilds(w http.ResponseWriter, r *http.Request)
	GetAgentDeployments(w http.ResponseWriter, r *http.Request)
	GetAgentEndpoints(w http.ResponseWriter, r *http.Request)
	GetBuild(w http.ResponseWriter, r *http.Request)
	GetAgentConfigurations(w http.ResponseWriter, r *http.Request)
	GetBuildLogs(w http.ResponseWriter, r *http.Request)
	GenerateName(w http.ResponseWriter, r *http.Request)
}

type agentController struct {
	agentService services.AgentManagerService
}

// NewAgentController returns a new AgentController instance.
func NewAgentController(agentService services.AgentManagerService) AgentController {
	return &agentController{
		agentService: agentService,
	}
}

func (c *agentController) GetAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)
	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("GetAgent: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	agent, err := c.agentService.GetAgent(ctx, userIdpId, orgName, projName, agentName)
	if err != nil {
		log.Error("GetAgent: failed to get agent", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get agent")
		return
	}

	agentResponse := utils.ConvertToAgentResponse(agent)
	utils.WriteSuccessResponse(w, http.StatusOK, agentResponse)
}

func (c *agentController) ListAgents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)
	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")

	// Validate required path parameters
	if orgName == "" || projName == "" {
		log.Error("ListAgents: missing required path parameters", "orgName", orgName, "projName", projName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	// Parse and validate pagination parameters
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 50 {
		log.Error("ListAgents: invalid limit parameter", "limit", limitStr)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter: must be between 1 and 50")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		log.Error("ListAgents: invalid offset parameter", "offset", offsetStr)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid offset parameter: must be 0 or greater")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	agents, total, err := c.agentService.ListAgents(ctx, userIdpId, orgName, projName, int32(limit), int32(offset))
	if err != nil {
		log.Error("ListAgents: failed to list agents", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to list agents")
		return
	}

	agentResponses := utils.ConvertToAgentListResponse(agents)
	response := &spec.AgentListResponse{
		Agents: agentResponses,
		Total:  total,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	utils.WriteSuccessResponse(w, http.StatusOK, response)
}

func (c *agentController) CreateAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")

	// Validate required path parameters
	if orgName == "" || projName == "" {
		log.Error("CreateAgent: missing required path parameters", "orgName", orgName, "projName", projName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	// Parse and validate request body
	var payload spec.CreateAgentRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error("CreateAgent: failed to decode request body", "error", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.ValidateAgentCreatePayload(payload); err != nil {
		log.Error("CreateAgent: invalid agent payload", "error", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := c.agentService.CreateAgent(ctx, userIdpId, orgName, projName, &payload)
	if err != nil {
		log.Error("CreateAgent: failed to create agent", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentAlreadyExists) {
			utils.WriteErrorResponse(w, http.StatusConflict, "Agent already exists")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create agent")
		return
	}
	response := &spec.AgentResponse{
		Name:         payload.Name,
		DisplayName:  payload.DisplayName,
		Description:  utils.StrPointerAsStr(payload.Description, ""),
		ProjectName:  projName,
		Provisioning: payload.Provisioning,
		CreatedAt:    time.Now(),
	}

	utils.WriteSuccessResponse(w, http.StatusAccepted, response)
}

func (c *agentController) DeleteAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("BuildAgent: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(r.Context())
	userIdpId := tokenClaims.Sub

	err := c.agentService.DeleteAgent(ctx, userIdpId, orgName, projName, agentName)
	if err != nil {
		log.Error("DeleteAgent: failed to delete agent", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to delete agent")
		return
	}
	utils.WriteSuccessResponse(w, http.StatusNoContent, "")
}

func (c *agentController) BuildAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("BuildAgent: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Parse query parameters
	commitId := r.URL.Query().Get("commitId")
	if commitId == "" {
		log.Debug("BuildAgent: commitId not provided, using latest commit")
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(r.Context())
	userIdpId := tokenClaims.Sub

	build, err := c.agentService.BuildAgent(ctx, userIdpId, orgName, projName, agentName, commitId)
	if err != nil {
		log.Error("BuildAgent: failed to build agent", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to build agent")
		return
	}
	utils.WriteSuccessResponse(w, http.StatusAccepted, build)
}

func (c *agentController) GetBuildLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")
	buildName := r.PathValue("buildName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" || buildName == "" {
		log.Error("GetBuildLogs: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName, "buildName", buildName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub
	buildLogs, err := c.agentService.GetBuildLogs(ctx, userIdpId, orgName, projName, agentName, buildName)
	if err != nil {
		log.Error("GetBuildLogs: failed to get build logs", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		if errors.Is(err, utils.ErrBuildNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Build not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get build logs")
		return
	}
	buildLogsResponse := utils.ConvertToBuildLogsResponse(*buildLogs)
	utils.WriteSuccessResponse(w, http.StatusOK, buildLogsResponse)
}

func (c *agentController) DeployAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("DeployAgent: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	// Parse and validate request body
	var payload spec.DeployAgentRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error("DeployAgent: failed to decode request body", "error", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if payload.ImageId == "" {
		log.Error("DeployAgent: imageId is required in request body")
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := c.agentService.DeployAgent(ctx, userIdpId, orgName, projName, agentName, &payload)
	if err != nil {
		log.Error("DeployAgent: failed to deploy agent", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to deploy agent")
		return
	}

	response := &spec.DeploymentResponse{
		AgentName:   agentName,
		ProjectName: projName,
		ImageId:     payload.ImageId,
		Environment: "Development",
	}
	utils.WriteSuccessResponse(w, http.StatusAccepted, response)
}

func (c *agentController) ListAgentBuilds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("ListAgentBuilds: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	// Parse and validate pagination parameters
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 50 {
		log.Error("ListAgentBuilds: invalid limit parameter", "limit", limitStr)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter: must be between 1 and 50")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		log.Error("ListAgentBuilds: invalid offset parameter", "offset", offsetStr)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid offset parameter: must be 0 or greater")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	builds, total, err := c.agentService.ListAgentBuilds(ctx, userIdpId, orgName, projName, agentName, int32(limit), int32(offset))
	if err != nil {
		log.Error("ListAgentBuilds: failed to list agent builds", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to list agent builds")
		return
	}

	buildResponses := utils.ConvertToBuildListResponse(builds)
	response := &spec.BuildsListResponse{
		Builds: buildResponses,
		Total:  total,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	utils.WriteSuccessResponse(w, http.StatusOK, response)
}

func (c *agentController) GenerateName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")

	// Validate required path parameters
	if orgName == "" {
		log.Error("CheckAgentNameAvailability: missing required path parameters", "orgName", orgName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	// Parse and validate request body
	var payload spec.ResourceNameRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error("GenerateName: failed to decode request body", "error", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := utils.ValidateResourceNameRequest(payload)
	if err != nil {
		log.Error("GenerateName: invalid resource name payload", "error", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid resource name payload")
		return
	}

	candidateName, err := c.agentService.GenerateName(ctx, userIdpId, orgName, payload)
	if err != nil {
		log.Error("GenerateAgentName: failed to generate agent name", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to check agent name availability")
		return
	}

	response := &spec.ResourceNameResponse{
		Name:         candidateName,
		DisplayName:  payload.DisplayName,
		ResourceType: payload.ResourceType,
	}
	utils.WriteSuccessResponse(w, http.StatusOK, response)
}

func (c *agentController) GetBuild(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")
	buildName := r.PathValue("buildName")

	if orgName == "" || projName == "" || agentName == "" || buildName == "" {
		log.Error("GetBuild: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName, "buildName", buildName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	build, err := c.agentService.GetBuild(ctx, userIdpId, orgName, projName, agentName, buildName)
	if err != nil {
		log.Error("GetBuild: failed to get build", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		if errors.Is(err, utils.ErrBuildNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Build not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get build")
		return
	}

	buildResponse := utils.ConvertToBuildDetailsResponse(build)
	utils.WriteSuccessResponse(w, http.StatusOK, buildResponse)
}

func (c *agentController) GetAgentDeployments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("GetAgentDeployments: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	deployments, err := c.agentService.GetAgentDeployments(ctx, userIdpId, orgName, projName, agentName)
	if err != nil {
		log.Error("GetAgentDeployments: failed to get deployments", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get deployments")
		return
	}

	deploymentResponses := utils.ConvertToDeploymentDetailsResponse(deployments)
	utils.WriteSuccessResponse(w, http.StatusOK, deploymentResponses)
}

func (c *agentController) GetAgentEndpoints(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")
	environment := r.URL.Query().Get("environment")
	if environment == "" {
		log.Error("GetAgentEndpoints: missing required query parameter 'environment'")
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required query parameter 'environment'")
		return
	}

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("GetAgentEndpoints: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	endpoints, err := c.agentService.GetAgentEndpoints(ctx, userIdpId, orgName, projName, agentName, environment)
	if err != nil {
		log.Error("GetAgentEndpoints: failed to get agent endpoints", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get agent endpoints")
		return
	}

	endpointResponses := utils.ConvertToAgentEndpointResponse(endpoints)
	utils.WriteSuccessResponse(w, http.StatusOK, endpointResponses)
}

func (c *agentController) GetAgentConfigurations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger(ctx)

	// Extract path parameters
	orgName := r.PathValue("orgName")
	projName := r.PathValue("projName")
	agentName := r.PathValue("agentName")

	// Validate required path parameters
	if orgName == "" || projName == "" || agentName == "" {
		log.Error("GetAgentConfigurations: missing required path parameters", "orgName", orgName, "projName", projName, "agentName", agentName)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required path parameters")
		return
	}

	environment := r.URL.Query().Get("environment")
	if environment == "" {
		log.Error("GetAgentConfigurations: missing required query parameter 'environment'")
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Missing required query parameter 'environment'")
		return
	}

	// Extract user info from JWT token
	tokenClaims := jwtassertion.GetTokenClaims(ctx)
	userIdpId := tokenClaims.Sub

	configurations, err := c.agentService.GetAgentConfigurations(ctx, userIdpId, orgName, projName, agentName, environment)
	if err != nil {
		log.Error("GetAgentConfigurations: failed to get configurations", "error", err)
		if errors.Is(err, utils.ErrOrganizationNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Organization not found")
			return
		}
		if errors.Is(err, utils.ErrProjectNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Project not found")
			return
		}
		if errors.Is(err, utils.ErrAgentNotFound) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Agent not found")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get configurations")
		return
	}

	// Convert configurations to response format
	configurationItems := make([]spec.ConfigurationItem, len(configurations))
	for i, config := range configurations {
		configurationItems[i] = spec.ConfigurationItem{
			Key:   config.Key,
			Value: config.Value,
		}
	}

	configurationsResponse := spec.ConfigurationResponse{
		ProjectName:    projName,
		AgentName:      agentName,
		Environment:    environment,
		Configurations: configurationItems,
	}

	utils.WriteSuccessResponse(w, http.StatusOK, configurationsResponse)
}
