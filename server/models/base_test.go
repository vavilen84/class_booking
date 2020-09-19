package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/test"
	"log"
)

var (
	testPilatesCapacity = 10

	TestPilatesClass = Class{
		Id:       "2b99f7e3-1e6a-47d5-839d-9fbff613bfbb",
		Name:     "Pilates",
		Capacity: &testPilatesCapacity,
	}
)

func PrepareTestDB() (db *sql.DB) {
	db = test.InitTestDb()
	test.DropAllTables(db)

	err := CreateMigrationsTableIfNotExists(db)
	if err != nil {
		log.Println(err)
	}

	err = MigrateUp(db)
	if err != nil {
		log.Println(err)
	}

	loadFixtures(db)
	return
}

func loadFixtures(db *sql.DB) {
	c := TestPilatesClass
	err := c.Insert(db)
	if err != nil {
		helpers.LogError(err)
	}
}
