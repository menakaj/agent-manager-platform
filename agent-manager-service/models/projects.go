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

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// API Response DTO
type ProjectResponse struct {
	Name               string    `json:"name"`
	OrgName            string    `json:"orgName"`
	DisplayName        string    `json:"displayName,omitempty"`
	Description        string    `json:"description,omitempty"`
	DeploymentPipeline string    `json:"deploymentPipeline,omitempty"`
	CreatedAt          time.Time `json:"createdAt"`
	Status             string    `json:"status,omitempty"`
}

// DB Model
type Project struct {
	ID                uuid.UUID      `gorm:"column:id;primaryKey" json:"projectID" binding:"required"`
	Name              string         `gorm:"column:name" json:"name" binding:"required"`
	OrgID             uuid.UUID      `gorm:"column:org_id" json:"orgID" binding:"required"`
	OpenChoreoProject string         `gorm:"column:open_choreo_project" json:"openChoreoProject,omitempty"`
	DisplayName       string         `gorm:"column:display_name" json:"displayName,omitempty"`
	Description       string         `gorm:"column:description" json:"description,omitempty"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
