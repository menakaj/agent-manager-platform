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

package apitestutils

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/db"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/models"
)

func CreateOrganization(t *testing.T, orgID uuid.UUID, userIdpID uuid.UUID, orgName string) models.Organization {
	org := &models.Organization{
		ID:                orgID,
		UserIdpId:         userIdpID,
		OrgName:           orgName,
		OpenChoreoOrgName: orgName,
		CreatedAt:         time.Now(),
	}
	err := db.DB(context.Background()).Create(org).Error
	require.NoError(t, err)
	str, _ := json.MarshalIndent(org, "", "  ")
	t.Logf("Created Organization: %s", str)
	return *org
}

func CreateProject(t *testing.T, projectID uuid.UUID, orgID uuid.UUID, projectName string) models.Project {
	project := &models.Project{
		ID:                projectID,
		OrgID:             orgID,
		Name:              projectName,
		CreatedAt:         time.Now(),
		OpenChoreoProject: projectName,
	}
	err := db.DB(context.Background()).Create(project).Error
	require.NoError(t, err)
	str, _ := json.MarshalIndent(project, "", "  ")
	t.Logf("Created Project: %s", str)
	return *project
}

func CreateAgent(t *testing.T, agentID uuid.UUID, orgID uuid.UUID, projectID uuid.UUID, agentName string) models.Agent {
	agent := &models.Agent{
		ID:          agentID,
		AgentType:   "internal",
		ProjectId:   projectID,
		OrgID:       orgID,
		Name:        agentName,
		DisplayName: agentName,
	}
	err := db.DB(context.Background()).Create(agent).Error
	require.NoError(t, err)
	str, _ := json.MarshalIndent(agent, "", "  ")
	t.Logf("Created Agent: %s", str)
	return *agent
}
