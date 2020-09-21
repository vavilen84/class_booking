package store

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/class_booking/helpers"
	"log"
	"os"
	"time"
)

var (
	testDb *sql.DB
)

func InitTestDB() {
	testDb = initTestDb()
}

func GetNewTestDBConn() (conn *sql.Conn, ctx context.Context) {
	ctx = GetDefaultDBContext()
	conn, err := testDb.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func initTestDb() (db *sql.DB) {

	sqlDriver := os.Getenv("SQL_DRIVER")
	testSqlDsn := os.Getenv("TEST_SQL_DSN")
	if (sqlDriver == "") || (testSqlDsn == "") {
		//if we run test outside docker using host machine sql server - we need to load .env vars
		err := godotenv.Load("../../.env")
		if err != nil {
			// TODO fix .env path outside of docker
			err = godotenv.Load("../.env")
			if err != nil {
				helpers.LogError(err)
			}
		}
		sqlDriver = os.Getenv("SQL_DRIVER")
		// use credentials without db in order to create test db
		testSqlDsn = os.Getenv("LOCALHOST_SQL_DSN")
		db, err = sql.Open(sqlDriver, testSqlDsn)
		if err != nil {
			panic("failed to connect sql server: " + err.Error())
		}
		ctx := GetDefaultDBContext()
		conn, err := db.Conn(ctx)
		if err != nil {
			helpers.LogError(err)
		}
		defer conn.Close()
		err = createTestDbIfNotExists(ctx, conn, os.Getenv("MYSQL_TEST_DATABASE"))
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

func createTestDbIfNotExists(ctx context.Context, conn *sql.Conn, dbName string) (err error) {
	_, err = conn.ExecContext(ctx, `CREATE DATABASE IF NOT EXISTS `+dbName)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
