package store

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/class_booking/helpers"
	"os"
)

var (
	db *sql.DB
)

func InitDB() {
	db = initDb()
}

func GetNewDBConn() (conn *sql.Conn, ctx context.Context) {
	// for endpoints testing, see main_test.go
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

func initDBForLocalhostAppRun() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		helpers.LogError(err)
	}
	sqlServerDsn := os.Getenv("LOCALHOST_SQL_DSN")
	mysqlDbName := os.Getenv("MYSQL_DATABASE")
	DbDsn := os.Getenv("LOCALHOST_DB_SQL_DSN")
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initDBForDockerMySql() *sql.DB {
	sqlServerDsn := os.Getenv("SQL_DSN")
	mysqlDbName := os.Getenv("MYSQL_DATABASE")
	DbDsn := os.Getenv("DB_SQL_DSN")
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initDBForHostMachineMySql() *sql.DB {
	sqlServerDsn := setHostAddress(os.Getenv("HOST_MACHINE_SQL_DSN"))
	mysqlDbName := setHostAddress(os.Getenv("MYSQL_DATABASE"))
	DbDsn := setHostAddress(os.Getenv("HOST_MACHINE_DB_SQL_DSN"))
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initDb() *sql.DB {
	docker := os.Getenv("DOCKER")
	if docker != "true" {
		return initDBForLocalhostAppRun()
	}
	dockerMySqlOnHostMachine := os.Getenv("DOCKER_MYSQL_HOST_MACHINE")
	if dockerMySqlOnHostMachine != "true" {
		return initDBForDockerMySql()
	}
	return initDBForHostMachineMySql()
}
