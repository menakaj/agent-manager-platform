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

package dbmigrations

import (
	"gorm.io/gorm"
)

// create table agents
var migration005 = migration{
	ID: 5,
	Migrate: func(db *gorm.DB) error {
		createTable := `CREATE TABLE agents
(
   id      UUID PRIMARY KEY,
   name          VARCHAR(100) NOT NULL,
   display_name  VARCHAR(100) NOT NULL,
   agent_type    VARCHAR(100) NOT NULL,
   description   TEXT,
   project_id    UUID NOT NULL,
   org_id        UUID NOT NULL,
   created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
   deleted_at    TIMESTAMPTZ,
   CONSTRAINT fk_agents_project_id FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
   CONSTRAINT fk_agents_org_id FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
   CONSTRAINT agent_type_enum check (agent_type in ('internal', 'external'))
)`

		createIndex := `CREATE UNIQUE INDEX uk_agents_name_project_org ON agents(name, project_id, org_id) WHERE deleted_at IS NULL`

		return db.Transaction(func(tx *gorm.DB) error {
			if err := runSQL(tx, createTable, createIndex); err != nil {
				return err
			}
			return nil
		})
	},
}
