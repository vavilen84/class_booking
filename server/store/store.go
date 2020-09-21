package store

import (
	"context"
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/helpers"
	"log"
	"os"
	"time"
)

func createDbIfNotExists(ctx context.Context, conn *sql.Conn, dbName string) (err error) {
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func GetDefaultDBContext() context.Context {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, constants.DefaultStoreTimeout)
	return ctx
}

func processInitDb(sqlServerDsn, mysqlDbName, DbDsn string) (db *sql.DB) {
	sqlDriver := os.Getenv("SQL_DRIVER")
	// use credentials without db in order to create db
	db, err := sql.Open(sqlDriver, sqlServerDsn)
	if err != nil {
		panic("failed to connect sql server: " + err.Error())
	}
	ctx := GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	defer conn.Close()
	err = createDbIfNotExists(ctx, conn, mysqlDbName)
	if err != nil {
		panic("failed to create test db: " + err.Error())
	}
	db, err = sql.Open(sqlDriver, DbDsn)
	if err != nil {
		panic("failed to database: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
