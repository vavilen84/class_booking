package models

//import (
//	"database/sql"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/class_booking/constants"
//	"github.com/vavilen84/class_booking/test"
//	"testing"
//)
//
//func TestVisitorTimetableItemValidateRequiredTag(t *testing.T) {
//	err := Validate(VisitorTimetableItem{})
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "Id"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "VisitorId"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "TimetableItemId"))
//}
//
//func TestVisitorTimetableItemValidateUuid4Tag(t *testing.T) {
//	v := VisitorTimetableItem{
//		Id:              "not valid uuid4",
//		VisitorId:       "not valid uuid4",
//		TimetableItemId: "not valid uuid4",
//	}
//	err := Validate(v)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "Id"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "VisitorId"))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "TimetableItemId"))
//}
//
//func TestVisitorTimetableItemBookingByVisitorAndTimetableItemExists(t *testing.T) {
//	db := PrepareTestDB()
//	m := test.TestVisitorTimetableItem
//	err, exists := m.BookingByVisitorAndTimetableItemExists(db)
//	assert.Nil(t, err)
//	assert.True(t, exists)
//
//	m = VisitorTimetableItem{
//		Id:              uuid.New().String(),
//		VisitorId:       uuid.New().String(),
//		TimetableItemId: uuid.New().String(),
//	}
//	err, exists = m.BookingByVisitorAndTimetableItemExists(db)
//	assert.Nil(t, err)
//	assert.False(t, exists)
//}
//
//func TestVisitorTimetableItemValidateTimetableItemExists(t *testing.T) {
//	db := PrepareTestDB()
//	m := test.TestVisitorTimetableItem
//	err := m.ValidateTimetableItemExists(db)
//	assert.Nil(t, err)
//
//	m = VisitorTimetableItem{
//		Id:              uuid.New().String(),
//		VisitorId:       uuid.New().String(),
//		TimetableItemId: uuid.New().String(),
//	}
//	err = m.ValidateTimetableItemExists(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.VisitorDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
//}
//
//func TestVisitorTimetableItemFindById(t *testing.T) {
//	db := PrepareTestDB()
//	m := VisitorTimetableItem{}
//	err := m.FindById(db, test.TestVisitorTimetableItem.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, test.TestVisitorTimetableItem.Id, m.Id)
//	assert.Equal(t, test.TestVisitorTimetableItem.VisitorId, m.VisitorId)
//	assert.Equal(t, test.TestVisitorTimetableItem.TimetableItemId, m.TimetableItemId)
//}
//
//func TestVisitorTimetableItemInsert(t *testing.T) {
//	db := PrepareTestDB()
//	id := uuid.New().String()
//	v := VisitorTimetableItem{
//		Id:              id,
//		VisitorId:       test.TestVisitor.Id,
//		TimetableItemId: test.TestTimetableItem.Id,
//	}
//	err := v.Insert(db)
//	assert.Nil(t, err)
//
//	v = VisitorTimetableItem{}
//	err = v.FindById(db, id)
//	assert.Nil(t, err)
//	assert.Equal(t, id, v.Id)
//	assert.Equal(t, test.TestVisitor.Id, v.VisitorId)
//	assert.Equal(t, test.TestTimetableItem.Id, v.TimetableItemId)
//}
//
//func TestVisitorTimetableItemInsertAlreadyExistingBooking(t *testing.T) {
//	db := PrepareTestDB()
//	v := test.TestVisitorTimetableItem
//	err := v.Insert(db)
//	assert.NotNil(t, err)
//
//	err, exists := v.BookingByVisitorAndTimetableItemExists(db)
//	assert.Nil(t, err)
//	assert.True(t, exists)
//}
//
//func TestVisitorTimetableItemInsertWithNotExistingVisitorAndTimetableItem(t *testing.T) {
//	db := PrepareTestDB()
//	m := VisitorTimetableItem{
//		Id:              uuid.New().String(),
//		VisitorId:       uuid.New().String(),
//		TimetableItemId: uuid.New().String(),
//	}
//	err := m.Insert(db)
//	assert.NotNil(t, err)
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.VisitorDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
//	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
//}
//
//func TestVisitorTimetableItemDelete(t *testing.T) {
//	db := PrepareTestDB()
//	v := test.TestVisitorTimetableItem
//	err := v.FindById(db, test.TestVisitorTimetableItem.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, test.TestVisitorTimetableItem.Id, v.Id)
//
//	v.Delete(db)
//
//	err = v.FindById(db, test.TestVisitorTimetableItem.Id)
//	assert.Equal(t, sql.ErrNoRows, err)
//}
