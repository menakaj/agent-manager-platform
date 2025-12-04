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
)

// DB Model
type Organization struct {
	ID                uuid.UUID `gorm:"column:id;primaryKey"`
	OrgName           string    `gorm:"column:org_name"`
	OpenChoreoOrgName string    `gorm:"column:open_choreo_org_name"`
	UserIdpId         uuid.UUID `gorm:"column:user_idp_id"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}

// API Response DTO
type OrganizationResponse struct {
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName,omitempty"`
	Description string    `json:"description,omitempty"`
	Namespace   string    `json:"namespace,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	Status      string    `json:"status,omitempty"`
}
