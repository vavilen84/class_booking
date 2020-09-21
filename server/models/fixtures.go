package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/class_booking/database"
	"github.com/vavilen84/class_booking/helpers"
	"time"
)

var (
	testPilatesCapacity = 10
	testNow             = time.Now()
	testTomorrow        = time.Now().Add(24 * time.Hour)

	TestPilatesClass = Class{
		Id:       "2b99f7e3-1e6a-47d5-839d-9fbff613bfbb",
		Name:     "Pilates",
		Capacity: &testPilatesCapacity,
	}
	TestYogaClass = Class{
		Id:       "2b99f7e3-1e6a-47d5-839d-9fbff613bfba",
		Name:     "Yoga",
		Capacity: &testPilatesCapacity,
	}
	TestVisitor = Visitor{
		Id:    "2b99f7e3-1e6a-47d5-839d-9fbff613bfbc",
		Email: "visitor@example.com",
	}
	TestVisitor2 = Visitor{
		Id:    "2b99f7e3-1e6a-47d5-839d-9fbff613bfbf",
		Email: "visitor2@example.com",
	}
	TestTimetableItem = TimetableItem{
		Id:      "044f4b01-3121-4e9d-b31b-07ba2aa4d544",
		ClassId: TestYogaClass.Id,
		Date:    &testNow,
	}
	TestTimetableItem2 = TimetableItem{
		Id:      "7c47e9dc-c451-4797-9888-083324b61352",
		ClassId: TestYogaClass.Id,
		Date:    &testTomorrow,
	}
	TestVisitorTimetableItem = VisitorTimetableItem{
		Id:              "1c00116e-5125-46b4-9b87-aa2767b85eaa",
		VisitorId:       TestVisitor2.Id,
		TimetableItemId: TestTimetableItem2.Id,
	}
)

func LoadFixtures(ctx context.Context, conn *sql.Conn) {
	c := TestPilatesClass
	err := database.Insert(ctx, conn, c)
	if err != nil {
		helpers.LogError(err)
	}

	y := TestYogaClass
	err = database.Insert(ctx, conn, y)
	if err != nil {
		helpers.LogError(err)
	}

	v := TestVisitor
	err = database.Insert(ctx, conn, v)
	if err != nil {
		helpers.LogError(err)
	}

	v2 := TestVisitor2
	err = database.Insert(ctx, conn, v2)
	if err != nil {
		helpers.LogError(err)
	}

	t := TestTimetableItem
	err = database.Insert(ctx, conn, t)
	if err != nil {
		helpers.LogError(err)
	}

	t2 := TestTimetableItem2
	err = database.Insert(ctx, conn, t2)
	if err != nil {
		helpers.LogError(err)
	}

	tv := TestVisitorTimetableItem
	err = database.Insert(ctx, conn, tv)
	if err != nil {
		helpers.LogError(err)
	}
}
