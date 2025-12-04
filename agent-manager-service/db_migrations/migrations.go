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
	"context"
	"fmt"
	"log/slog"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/db"
)

var migrateOptions = &gormigrate.Options{
	TableName:                 "migration_history",
	IDColumnName:              "id",
	IDColumnSize:              255,
	UseTransaction:            true, // Controls whether migrations run within database transactions
	ValidateUnknownMigrations: true, // Controls validation of migrations that exist in the database but not in the code
}

type migration struct {
	ID      int32
	Migrate gormigrate.MigrateFunc
}

func Migrate() error {
	dbConn := db.DB(context.Background())

	successCount := 0
	var list []*gormigrate.Migration
	for _, m := range migrations {
		m := m
		id := generateIdStr(m.ID)
		list = append(list, &gormigrate.Migration{
			ID: id,
			Migrate: func(g *gorm.DB) error {
				slog.Info("dbmigrations:applying migration", "id", id)
				if err := m.Migrate(g); err != nil {
					return err
				}
				successCount++
				slog.Info("dbmigrations:migration applied successfully", "id", id)
				return nil
			},
		})
	}
	latestId := generateIdStr(latestVersion)
	slog.Info("dbmigrations:starting migration", "latest", latestId)
	m := gormigrate.New(dbConn, migrateOptions, list)
	if err := m.MigrateTo(latestId); err != nil {
		return err
	}
	slog.Info("dbmigrations:migration completed", "latest", latestId, "successCount", successCount)
	return nil
}

func generateIdStr(id int32) string {
	return fmt.Sprintf("%04d", id)
}
