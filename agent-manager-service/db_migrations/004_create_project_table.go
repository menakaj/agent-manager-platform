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

// This is a placeholder migration file

import (
	"gorm.io/gorm"
)

// create table projects
var migration004 = migration{
	ID: 4,
	Migrate: func(db *gorm.DB) error {
		createTable := `CREATE TABLE projects
(
    id   UUID PRIMARY KEY,
	name         VARCHAR(100) NOT NULL,
	org_id       UUID NOT NULL,
	open_choreo_project VARCHAR(100) NOT NULL,
	display_name VARCHAR(100),
	description  TEXT,
	created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at   TIMESTAMPTZ,
	CONSTRAINT fk_projects_org_id FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
	CONSTRAINT chk_projects_name_open_choreo_match CHECK (name = open_choreo_project)
)`

		createIndex := `CREATE UNIQUE INDEX uk_projects_name_org_id ON projects(name, org_id) WHERE deleted_at IS NULL`

		return db.Transaction(func(tx *gorm.DB) error {
			return runSQL(tx, createTable, createIndex)
		})
	},
}
