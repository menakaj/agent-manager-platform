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

type OrganizationRepository interface {
	GetOrganizationsByUserIdpID(ctx context.Context, userIdpID uuid.UUID) ([]models.Organization, error)
	CreateOrganization(ctx context.Context, organization *models.Organization) error
	GetOrganizationByOrgName(ctx context.Context, userIdpID uuid.UUID, orgName string) (*models.Organization, error)
}

type organizationRepository struct{}

func NewOrganizationRepository() OrganizationRepository {
	return &organizationRepository{}
}

func (r *organizationRepository) GetOrganizationsByUserIdpID(ctx context.Context, userIdpID uuid.UUID) ([]models.Organization, error) {
	var orgs []models.Organization
	if err := db.DB(ctx).Where("user_idp_id = ?", userIdpID).Find(&orgs).Error; err != nil {
		return nil, fmt.Errorf("organizationRepository.GetOrganizationsByUserIdpID: %w", err)
	}
	return orgs, nil
}

func (r *organizationRepository) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	if err := db.DB(ctx).Create(organization).Error; err != nil {
		return fmt.Errorf("organizationRepository.CreateOrganization: %w", err)
	}
	return nil
}

func (r *organizationRepository) GetOrganizationByOrgName(ctx context.Context, userIdpID uuid.UUID, orgName string) (*models.Organization, error) {
	var org models.Organization
	if err := db.DB(ctx).Where("user_idp_id = ? AND org_name = ?", userIdpID, orgName).First(&org).Error; err != nil {
		return nil, fmt.Errorf("organizationRepository.GetOrganizationByOrgName: %w", err)
	}
	return &org, nil
}
