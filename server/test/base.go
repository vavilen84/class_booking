package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/helpers"
	"log"
	"os"
	"time"
)

func InitTestDb() (db *sql.DB) {

	sqlDriver := os.Getenv("SQL_DRIVER")
	testSqlDsn := os.Getenv("TEST_SQL_DSN")
	if (sqlDriver == "") || (testSqlDsn == "") {
		//if we run test outside docker using host machine sql server - we need to load .env vars
		err := godotenv.Load("../../.env")
		if err != nil {
			helpers.LogError(err)
		}
		sqlDriver = os.Getenv("SQL_DRIVER")
		// use credentials without db in order to create test db
		testSqlDsn = os.Getenv("LOCALHOST_SQL_DSN")
		db, err = sql.Open(sqlDriver, testSqlDsn)
		if err != nil {
			panic("failed to connect sql server: " + err.Error())
		}
		err = createTestDbIfNotExists(db, os.Getenv("MYSQL_TEST_DATABASE"))
		if err != nil {
			panic("failed to create test db: " + err.Error())
		}
		testSqlDsn = os.Getenv("LOCALHOST_TEST_DB_SQL_DSN")
	}

	db, err := sql.Open(sqlDriver, testSqlDsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func createTestDbIfNotExists(db *sql.DB, dbName string) (err error) {
	query := `CREATE DATABASE IF NOT EXISTS ` + dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func DropAllTables(db *sql.DB) {
	tables := []string{
		constants.MigrationsTableName,
		constants.VisitorTimetableItemTableName,
		constants.VisitorTableName,
		constants.TimetableItemTableName,
		constants.ClassTableName,
	}
	for i := 0; i < len(tables); i++ {
		_, err := db.Exec("DROP TABLE IF EXISTS " + tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}
