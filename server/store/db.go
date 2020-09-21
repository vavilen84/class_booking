package store

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/class_booking/helpers"
	"log"
	"os"
	"time"
)

var (
	db *sql.DB
)

func InitDB() {
	db = initDb()
}

func GetNewDBConn() (conn *sql.Conn, ctx context.Context) {
	if os.Getenv("APP_ENV") == "test" {
		return GetNewTestDBConn()
	}
	ctx = GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func createDbIfNotExists(ctx context.Context, conn *sql.Conn, dbName string) (err error) {
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS ?", dbName)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func initDb() (db *sql.DB) {
	sqlDriver := os.Getenv("SQL_DRIVER")
	sqlDsn := os.Getenv("SQL_DSN")
	if (sqlDriver == "") || (sqlDsn == "") {
		//if we run test outside docker using host machine sql server - we need to load .env vars
		err := godotenv.Load("../.env")
		if err != nil {
			helpers.LogError(err)
		}
		sqlDriver = os.Getenv("SQL_DRIVER")
		// use credentials without db in order to create db
		sqlDsn = os.Getenv("LOCALHOST_SQL_DSN")
		db, err = sql.Open(sqlDriver, sqlDsn)
		if err != nil {
			panic("failed to connect sql server: " + err.Error())
		}
		ctx := GetDefaultDBContext()
		conn, err := db.Conn(ctx)
		if err != nil {
			helpers.LogError(err)
		}
		defer conn.Close()
		err = createDbIfNotExists(ctx, conn, os.Getenv("MYSQL_DATABASE"))
		if err != nil {
			panic("failed to create test db: " + err.Error())
		}
		sqlDsn = os.Getenv("LOCALHOST_DB_SQL_DSN")
	}

	db, err := sql.Open(sqlDriver, sqlDsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
