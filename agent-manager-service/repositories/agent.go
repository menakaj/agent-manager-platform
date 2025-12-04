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

package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/db"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/models"
)

type AgentRepository interface {
	ListAgents(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID) ([]*models.Agent, error)
	GetAgentByName(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) (*models.Agent, error)
	CreateAgent(ctx context.Context, agent *models.Agent) error
	DeleteAgentByName(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) error
	UpdateAgentTimestamp(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) error
}

type agentRepository struct{}

func NewAgentRepository() AgentRepository {
	return &agentRepository{}
}

func (r *agentRepository) ListAgents(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID) ([]*models.Agent, error) {
	var agents []*models.Agent
	if err := db.DB(ctx).Where("org_id = ? AND project_id = ?", orgId, projectId).Find(&agents).Error; err != nil {
		return nil, fmt.Errorf("agentRepository.ListAgents: %w", err)
	}

	return agents, nil
}

func (r *agentRepository) GetAgentByName(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) (*models.Agent, error) {
	var agent models.Agent
	if err := db.DB(ctx).Where("org_id = ? AND project_id = ? AND name = ?", orgId, projectId, agentName).First(&agent).Error; err != nil {
		return nil, fmt.Errorf("agentRepository.GetAgentByName: %w", err)
	}
	return &agent, nil
}

func (r *agentRepository) CreateAgent(ctx context.Context, agent *models.Agent) error {
	if err := db.DB(ctx).Create(agent).Error; err != nil {
		return fmt.Errorf("agentRepository.CreateAgent: %w", err)
	}
	return nil
}

func (r *agentRepository) DeleteAgentByName(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) error {
	if err := db.DB(ctx).Where("org_id = ? AND project_id = ? AND name = ?", orgId, projectId, agentName).Delete(&models.Agent{}).Error; err != nil {
		return fmt.Errorf("agentRepository.DeleteAgentByName: %w", err)
	}
	return nil
}

func (r *agentRepository) UpdateAgentTimestamp(ctx context.Context, orgId uuid.UUID, projectId uuid.UUID, agentName string) error {
	if err := db.DB(ctx).Model(&models.Agent{}).
		Where("org_id = ? AND project_id = ? AND name = ?", orgId, projectId, agentName).
		Update("updated_at", "NOW()").Error; err != nil {
		return fmt.Errorf("agentRepository.UpdateAgentTimestamp: %w", err)
	}
	return nil
}
