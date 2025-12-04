package dbmigrations

import (
	"gorm.io/gorm"
)

var migration006 = migration{
	ID: 6,
	Migrate: func(db *gorm.DB) error {
		insertDefaultOrg := `INSERT INTO organizations 
(id, open_choreo_org_name, user_idp_id, org_name) 
VALUES 
('af779290-c22d-4100-aefd-484d81fff60e', 'default', '8f307351-25c5-4fc6-85e0-f51c2d458f06', 'default');
`
		insertDefaultProject := `INSERT INTO projects 
(id, name, open_choreo_project, org_id, display_name) 
VALUES 
('9c2c0915-7c33-4cb8-8ceb-030aff811f8d', 'default', 'default','af779290-c22d-4100-aefd-484d81fff60e', 'Default');
`
		return db.Transaction(func(tx *gorm.DB) error {
			return runSQL(tx, insertDefaultOrg, insertDefaultProject)
		})
	},
}
