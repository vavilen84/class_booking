package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/class_booking/migrate"
	"log"
	"os"
	"time"
)

func main() {
	db, err := sql.Open(os.Getenv("SQL_DRIVER"), os.Getenv("SQL_DSN"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = migrate.CreateMigrationsTableIfNotExists(db)
	if err != nil {
		log.Println(err)
	}

	err = migrate.MigrateUp(db)
	if err != nil {
		log.Println(err)
	}
}
