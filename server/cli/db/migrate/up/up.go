package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/class_booking/models"
	"github.com/vavilen84/class_booking/store"
	"log"
)

func main() {
	store.InitDB()
	conn, ctx := store.GetNewDBConn()
	defer conn.Close()

	err := models.CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = models.MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}
}
