package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/migrate"
	"log"
	"os"
)

func initTestDB() (db *sql.DB) {
	db, err := sql.Open(os.Getenv("SQL_DRIVER"), os.Getenv("TEST_SQL_DSN"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func createTestDbIfNotExists(db *sql.DB) (err error) {
	query := `
		CREATE DATABASE IF NOT EXISTS class_booking_test` + constants.MigrationsTableName + `
		(
			version integer PRIMARY KEY,
			filename text NOT NULL,
			created_at integer NOT NULL
		)
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func clearDb(db *sql.DB) {
	dropAllTables(db)

	err := migrate.CreateMigrationsTableIfNotExists(db)
	if err != nil {
		log.Println(err)
	}

	err = migrate.MigrateUp(db)
	if err != nil {
		log.Println(err)
	}
}

func dropAllTables(db *sql.DB) {
	tables := []string{
		constants.MigrationsTableName,
		constants.ClassTableName,
	}
	for i := 0; i < len(tables); i++ {
		_, err := db.Exec("DROP TABLE IF EXISTS " + tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}
