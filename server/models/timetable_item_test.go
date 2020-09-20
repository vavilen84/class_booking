package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/store"
	"testing"
	"time"
)

func TestTimetableItemValidateRequiredTag(t *testing.T) {
	err := Validate(TimetableItem{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "ClassId"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "Date"))
}

func TestTimetableItemValidateUuid4Tag(t *testing.T) {
	v := TimetableItem{
		Id:      "not valid uuid4",
		ClassId: "not valid uuid4",
	}
	err := Validate(v)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.TimetableItemStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.TimetableItemStructName, "ClassId"))
}

func TestTimetableItemFindByDate(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := TimetableItem{}
	err := m.FindByDate(ctx, conn, TestTimetableItem.Date)
	assert.Nil(t, err)
	assert.Equal(t, TestTimetableItem.Id, m.Id)
	assert.Equal(t, TestTimetableItem.ClassId, m.ClassId)
	assert.Equal(t, TestTimetableItem.Date.Format(constants.DateFormat), m.Date.Format(constants.DateFormat))
}

func TestTimetableItemFindById(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := TimetableItem{}
	err := m.FindById(ctx, conn, TestTimetableItem.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestTimetableItem.Id, m.Id)
	assert.Equal(t, TestTimetableItem.ClassId, m.ClassId)
	assert.Equal(t, TestTimetableItem.Date.Format(constants.DateFormat), m.Date.Format(constants.DateFormat))
}

func TestTimetableItemInsert(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	d := testNow.Add(48 * time.Hour)
	v := TimetableItem{
		Id:      id,
		ClassId: TestYogaClass.Id,
		Date:    &d,
	}
	err := v.Insert(ctx, conn)
	assert.Nil(t, err)

	v = TimetableItem{}
	err = v.FindById(ctx, conn, id)
	assert.Nil(t, err)
	assert.Equal(t, id, v.Id)
	assert.Equal(t, TestYogaClass.Id, v.ClassId)
	assert.Equal(t, d.Format(constants.DateFormat), v.Date.Format(constants.DateFormat))
}

func TestTimetableItemValidateDate(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	v := TimetableItem{
		Id:      id,
		ClassId: TestYogaClass.Id,
		Date:    TestTimetableItem.Date,
	}
	err := v.ValidateDate(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
}

func TestTimetableItemValidateClassExists(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	v := TimetableItem{
		Id:      id,
		ClassId: "not_existing_class_id",
		Date:    TestTimetableItem.Date,
	}
	err := v.ValidateClassExists(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.ClassDoesNotExistErrorMsg, constants.TimetableItemStructName))
}

func TestTimetableItemInsertWithAlreadyRegisteredDate(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	v := TimetableItem{
		Id:      id,
		ClassId: TestYogaClass.Id,
		Date:    TestTimetableItem.Date,
	}
	err := v.Insert(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
}

func TestTimetableItemInsertWithNotExistingClass(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	d := time.Now().Add(72 * time.Hour)
	v := TimetableItem{
		Id:      id,
		ClassId: uuid.New().String(),
		Date:    &d,
	}
	err := v.Insert(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.ClassDoesNotExistErrorMsg, constants.TimetableItemStructName))
}

func TestTimetableItemDelete(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	v := TimetableItem{}
	err := v.FindById(ctx, conn, TestTimetableItem.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestTimetableItem.Id, v.Id)

	err = v.Delete(ctx, conn)
	assert.Nil(t, err)

	err = v.FindById(ctx, conn, TestTimetableItem.Id)
	assert.Equal(t, sql.ErrNoRows, err)
}
