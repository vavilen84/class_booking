package models

//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/class_booking/constants"
//	"github.com/vavilen84/class_booking/test"
//	"testing"
//	"time"
//)
//
//func TestTimetableItemValidateRequiredTag(t *testing.T) {
//	err := Validate(TimetableItem{})
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "Id"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "ClassId"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.TimetableItemStructName, "Date"))
//}
//
//func TestTimetableItemValidateUuid4Tag(t *testing.T) {
//	v := TimetableItem{
//		Id:      "not valid uuid4",
//		ClassId: "not valid uuid4",
//	}
//	err := Validate(v)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.TimetableItemStructName, "Id"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.TimetableItemStructName, "ClassId"))
//}
//
//func TestTimetableItemFindByDate(t *testing.T) {
//	db := PrepareTestDB()
//	m := TimetableItem{}
//	err := m.FindByDate(db, test.TestTimetableItem.Date)
//	assert.Nil(t, err)
//	assert.Equal(t, test.TestTimetableItem.Id, m.Id)
//	assert.Equal(t, test.TestTimetableItem.ClassId, m.ClassId)
//	assert.Equal(t, test.TestTimetableItem.Date.Format(constants.DateFormat), m.Date.Format(constants.DateFormat))
//}
//
//func TestTimetableItemFindById(t *testing.T) {
//	db := PrepareTestDB()
//	m := TimetableItem{}
//	err := m.FindById(db, test.TestTimetableItem.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, test.TestTimetableItem.Id, m.Id)
//	assert.Equal(t, test.TestTimetableItem.ClassId, m.ClassId)
//	assert.Equal(t, test.TestTimetableItem.Date.Format(constants.DateFormat), m.Date.Format(constants.DateFormat))
//}
//
//func TestTimetableItemInsert(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	d := test.testNow.Add(48 * time.Hour)
//	v := TimetableItem{
//		Id:      id,
//		ClassId: test.TestYogaClass.Id,
//		Date:    &d,
//	}
//	err := v.Insert(db)
//	assert.Nil(t, err)
//
//	v = TimetableItem{}
//	err = v.FindById(db, id)
//	assert.Nil(t, err)
//	assert.Equal(t, id, v.Id)
//	assert.Equal(t, test.TestYogaClass.Id, v.ClassId)
//	assert.Equal(t, d.Format(constants.DateFormat), v.Date.Format(constants.DateFormat))
//}
//
//func TestTimetableItemValidateDate(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	v := TimetableItem{
//		Id:      id,
//		ClassId: test.TestYogaClass.Id,
//		Date:    test.TestTimetableItem.Date,
//	}
//	err := v.ValidateDate(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
//}
//
//func TestTimetableItemValidateClassExists(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	v := TimetableItem{
//		Id:      id,
//		ClassId: "not_existing_class_id",
//		Date:    test.TestTimetableItem.Date,
//	}
//	err := v.ValidateClassExists(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.ClassDoesNotExistErrorMsg, constants.TimetableItemStructName))
//}
//
//func TestTimetableItemInsertWithAlreadyRegisteredDate(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	v := TimetableItem{
//		Id:      id,
//		ClassId: test.TestYogaClass.Id,
//		Date:    test.TestTimetableItem.Date,
//	}
//	err := v.Insert(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
//}
//
//func TestTimetableItemInsertWithNotExistingClass(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	d := time.Now().Add(72 * time.Hour)
//	v := TimetableItem{
//		Id:      id,
//		ClassId: uuid.New().String(),
//		Date:    &d,
//	}
//	err := v.Insert(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.ClassDoesNotExistErrorMsg, constants.TimetableItemStructName))
//}
//
//func TestTimetableItemDelete(t *testing.T) {
//	db := PrepareTestDB()
//	v := TimetableItem{}
//	err := v.FindById(db, test.TestTimetableItem.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, test.TestTimetableItem.Id, v.Id)
//
//	v.Delete(db)
//
//	err = v.FindById(db, test.TestTimetableItem.Id)
//	assert.Equal(t, sql.ErrNoRows, err)
//}
