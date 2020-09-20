package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/test"
	"log"
	"math/rand"
	"time"
)

var (
	testPilatesCapacity = 10

	testNow        = time.Now()
	testSeededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	testCharset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	TestPilatesClass = Class{
		Id:       "2b99f7e3-1e6a-47d5-839d-9fbff613bfbb",
		Name:     "Pilates",
		Capacity: &testPilatesCapacity,
	}
	TestYogaClass = Class{
		Id:       "2b99f7e3-1e6a-47d5-839d-9fbff613bfba",
		Name:     "Pilates",
		Capacity: &testPilatesCapacity,
	}
	TestVisitor = Visitor{
		Id:    "2b99f7e3-1e6a-47d5-839d-9fbff613bfbc",
		Email: "visitor@example.com",
	}
	TestTimetableItem = TimetableItem{
		Id:      "2b99f7e3-1e6a-47d5-839d-9fbff613bfby",
		ClassId: TestYogaClass.Id,
		Date:    &testNow,
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
	err := Insert(db, c)
	if err != nil {
		helpers.LogError(err)
	}

	y := TestYogaClass
	err = Insert(db, y)
	if err != nil {
		helpers.LogError(err)
	}

	v := TestVisitor
	err = Insert(db, v)
	if err != nil {
		helpers.LogError(err)
	}

	t := TestTimetableItem
	err = Insert(db, t)
	if err != nil {
		helpers.LogError(err)
	}
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = testCharset[testSeededRand.Intn(len(testCharset))]
	}
	return string(b)
}
