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
type AgentResponse struct {
	Name         string       `json:"name"`
	DisplayName  string       `json:"displayName,omitempty"`
	Description  string       `json:"description,omitempty"`
	ProjectName  string       `json:"projectName"`
	CreatedAt    time.Time    `json:"createdAt"`
	Status       string       `json:"status,omitempty"`
	Provisioning Provisioning `json:"provisioning,omitempty"`
}

type Provisioning struct {
	Type       string     `json:"type"`
	Repository Repository `json:"repository,omitempty"`
}

type Repository struct {
	Url     string `json:"url"`
	AppPath string `json:"appPath"`
	Branch  string `json:"branch"`
}

// DB Model
type Agent struct {
	ID          uuid.UUID      `gorm:"column:id;primaryKey"`
	AgentType   string         `gorm:"column:agent_type"`
	Name        string         `gorm:"column:name"`
	DisplayName string         `gorm:"column:display_name"`
	Description string         `gorm:"column:description"`
	ProjectId   uuid.UUID      `gorm:"column:project_id"`
	OrgID       uuid.UUID      `gorm:"column:org_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
