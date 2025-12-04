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

// add org_name column to organizations table
var migration003 = migration{
	ID: 3,
	Migrate: func(db *gorm.DB) error {
		return db.Transaction(func(tx *gorm.DB) error {
			// Add the column without NOT NULL constraint first
			addColumn := `ALTER TABLE organizations ADD COLUMN org_name VARCHAR(100);`
			if err := runSQL(tx, addColumn); err != nil {
				return err
			}

			// Set default values from open_choreo_org_name
			updateValues := `UPDATE organizations SET org_name = open_choreo_org_name WHERE org_name IS NULL;`
			if err := runSQL(tx, updateValues); err != nil {
				return err
			}

			// Now add NOT NULL and UNIQUE constraints
			addConstraints := `ALTER TABLE organizations 
				ALTER COLUMN org_name SET NOT NULL,
				ADD CONSTRAINT organizations_org_name_unique UNIQUE (org_name),
				ADD CONSTRAINT chk_organizations_name_match CHECK (org_name = open_choreo_org_name);`
			return runSQL(tx, addConstraints)
		})
	},
}
