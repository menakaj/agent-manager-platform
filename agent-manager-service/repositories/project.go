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

type ProjectRepository interface {
	ListProjects(ctx context.Context, orgId uuid.UUID) ([]models.Project, error)
	GetProjectByName(ctx context.Context, orgId uuid.UUID, projectName string) (*models.Project, error)
	CreateProject(ctx context.Context, project *models.Project) error
}

type projectRepository struct{}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{}
}

func (r *projectRepository) ListProjects(ctx context.Context, orgId uuid.UUID) ([]models.Project, error) {
	var projects []models.Project
	if err := db.DB(ctx).Where("org_id = ?", orgId).Find(&projects).Error; err != nil {
		return nil, fmt.Errorf("projectRepository.GetProjects: %w", err)
	}

	return projects, nil
}

func (r *projectRepository) GetProjectByName(ctx context.Context, orgId uuid.UUID, projectName string) (*models.Project, error) {
	var project models.Project
	if err := db.DB(ctx).Where("org_id = ? AND name = ?", orgId, projectName).First(&project).Error; err != nil {
		return nil, fmt.Errorf("projectRepository.GetProjectByName: %w", err)
	}
	return &project, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, project *models.Project) error {
	if err := db.DB(ctx).Create(project).Error; err != nil {
		return fmt.Errorf("projectRepository.CreateProject: %w", err)
	}
	return nil
}
