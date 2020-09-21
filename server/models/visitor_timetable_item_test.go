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

func TestVisitorTimetableItemValidateRequiredTag(t *testing.T) {
	err := Validate(VisitorTimetableItem{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "VisitorId"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorTimetableItemStructName, "TimetableItemId"))
}

func TestVisitorTimetableItemValidateUuid4Tag(t *testing.T) {
	v := VisitorTimetableItem{
		Id:              "not valid uuid4",
		VisitorId:       "not valid uuid4",
		TimetableItemId: "not valid uuid4",
	}
	err := Validate(v)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "VisitorId"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorTimetableItemStructName, "TimetableItemId"))
}

func TestVisitorTimetableItemBookingByVisitorAndTimetableItemExists(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	m := TestVisitorTimetableItem
	err, exists := m.BookingByVisitorAndTimetableItemExists(ctx, conn)
	assert.Nil(t, err)
	assert.True(t, exists)

	m = VisitorTimetableItem{
		Id:              uuid.New().String(),
		VisitorId:       uuid.New().String(),
		TimetableItemId: uuid.New().String(),
	}
	err, exists = m.BookingByVisitorAndTimetableItemExists(ctx, conn)
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestVisitorTimetableItemValidateTimetableItemExists(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	m := TestVisitorTimetableItem
	err := m.ValidateTimetableItemExists(ctx, conn)
	assert.Nil(t, err)

	m = VisitorTimetableItem{
		Id:              uuid.New().String(),
		VisitorId:       uuid.New().String(),
		TimetableItemId: uuid.New().String(),
	}
	err = m.ValidateTimetableItemExists(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.VisitorDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
}

func TestVisitorTimetableItemFindById(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	m := VisitorTimetableItem{}
	err := m.FindById(ctx, conn, TestVisitorTimetableItem.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitorTimetableItem.Id, m.Id)
	assert.Equal(t, TestVisitorTimetableItem.VisitorId, m.VisitorId)
	assert.Equal(t, TestVisitorTimetableItem.TimetableItemId, m.TimetableItemId)
}

func TestVisitorTimetableItemFindByVisitorEmail(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	m := VisitorTimetableItem{}
	err := m.FindByVisitorEmail(ctx, conn, TestVisitor2.Email)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitorTimetableItem.VisitorId, TestVisitor2.Id)
}

func TestVisitorTimetableItemInsert(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	id := uuid.New().String()
	v := VisitorTimetableItem{
		Id:              id,
		VisitorId:       TestVisitor.Id,
		TimetableItemId: TestTimetableItem.Id,
	}
	err := v.Insert(ctx, conn)
	assert.Nil(t, err)

	v = VisitorTimetableItem{}
	err = v.FindById(ctx, conn, id)
	assert.Nil(t, err)
	assert.Equal(t, id, v.Id)
	assert.Equal(t, TestVisitor.Id, v.VisitorId)
	assert.Equal(t, TestTimetableItem.Id, v.TimetableItemId)
}

func TestVisitorTimetableItemInsertAlreadyExistingBooking(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	v := TestVisitorTimetableItem
	err := v.Insert(ctx, conn)
	assert.NotNil(t, err)

	err, exists := v.BookingByVisitorAndTimetableItemExists(ctx, conn)
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestVisitorTimetableItemInsertWithNotExistingVisitorAndTimetableItem(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	m := VisitorTimetableItem{
		Id:              uuid.New().String(),
		VisitorId:       uuid.New().String(),
		TimetableItemId: uuid.New().String(),
	}
	err := m.Insert(ctx, conn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.VisitorDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
}

func TestVisitorTimetableItemDelete(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	v := TestVisitorTimetableItem
	err := v.FindById(ctx, conn, TestVisitorTimetableItem.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitorTimetableItem.Id, v.Id)

	err = v.Delete(ctx, conn)
	assert.Nil(t, err)

	err = v.FindById(ctx, conn, TestVisitorTimetableItem.Id)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestVisitorTimetableItemValidateAPIBookings(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	ab := APIBookings{}
	vti := VisitorTimetableItem{}
	err := vti.ValidateAPIBookings(ctx, conn, ab)
	assert.NotNil(t, err)

	ab = APIBookings{
		Email: "not_existing_email",
		Date:  TestTimetableItem.Date,
	}
	err = vti.ValidateAPIBookings(ctx, conn, ab)
	assert.NotNil(t, err)

	notRegisteredDate := time.Now().AddDate(3, 0, 0)
	ab = APIBookings{
		Email: TestVisitor.Email,
		Date:  &notRegisteredDate,
	}
	err = vti.ValidateAPIBookings(ctx, conn, ab)
	assert.NotNil(t, err)

	ab = APIBookings{
		Email: TestVisitor.Email,
		Date:  TestTimetableItem.Date,
	}
	err = vti.ValidateAPIBookings(ctx, conn, ab)
	assert.Nil(t, err)
}
