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

// create table agent
var migration002 = migration{
	ID: 2,
	Migrate: func(db *gorm.DB) error {
		createTable := `CREATE TABLE organizations
(
    id              UUID PRIMARY KEY,
    open_choreo_org_name VARCHAR(100) NOT NULL UNIQUE,
    user_idp_id         UUID NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

		return db.Transaction(func(tx *gorm.DB) error {
			return runSQL(tx, createTable)
		})
	},
}
