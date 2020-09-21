package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/store"
	"log"
)

func init() {
	store.InitTestDB()
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB(ctx context.Context, conn *sql.Conn) {
	dropAllTablesFromTestDB(ctx, conn)
	err := CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	LoadFixtures(ctx, conn)
	return
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func dropAllTablesFromTestDB(ctx context.Context, conn *sql.Conn) {
	tables := []string{
		constants.MigrationsTableName,
		constants.VisitorTimetableItemTableName,
		constants.VisitorTableName,
		constants.TimetableItemTableName,
		constants.ClassTableName,
	}
	for i := 0; i < len(tables); i++ {
		_, err := conn.ExecContext(ctx, "DROP TABLE IF EXISTS "+tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}
