package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/containers"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/store"
	"testing"
	"time"
)

func TestClassValidateRequiredTag(t *testing.T) {
	err := Validate(Class{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.ClassStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.ClassStructName, "Name"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.ClassStructName, "Capacity"))
}

func TestClassValidateMinValueTag(t *testing.T) {
	notValidCapacity := 0
	c := Class{
		Name:     "n",
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, constants.ClassStructName, "Name", "2"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, constants.ClassStructName, "Capacity", "1"))
}

func TestClassValidateUuid4Tag(t *testing.T) {
	c := Class{
		Id: "not valid uuid4",
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.ClassStructName, "Id"))
}

func TestClassValidateMaxValueTag(t *testing.T) {
	notValidCapacity := 51
	c := Class{
		Name:     helpers.GenerateRandomString(256),
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, constants.ClassStructName, "Name", "255"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, constants.ClassStructName, "Capacity", "50"))
}

func TestClassInsert(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	capacity := 10
	id := uuid.New().String()
	name := "Crossfit"
	c := Class{
		Id:       id,
		Name:     name,
		Capacity: &capacity,
	}
	err := c.Insert(ctx, conn)
	assert.Nil(t, err)

	c = Class{}
	err = c.FindById(ctx, conn, id)
	assert.Nil(t, err)
	assert.Equal(t, c.Id, id)
	assert.Equal(t, c.Name, name)
	assert.Equal(t, c.Capacity, &capacity)
}

func TestClassFindById(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := Class{}
	err := m.FindById(ctx, conn, TestPilatesClass.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, TestPilatesClass.Name, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)
}

func TestClassUpdate(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := Class{}
	err := m.FindById(ctx, conn, TestPilatesClass.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, TestPilatesClass.Name, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)

	updatedName := "Updated Name"
	m.Name = updatedName
	err = m.Update(ctx, conn)
	assert.Nil(t, err)
	assert.Equal(t, updatedName, m.Name)

	err = m.FindById(ctx, conn, TestPilatesClass.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, updatedName, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)
}

func TestClassDelete(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := Class{}
	err := m.FindById(ctx, conn, TestPilatesClass.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, TestPilatesClass.Name, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)

	err = m.Delete(ctx, conn)
	assert.Nil(t, err)

	err = m.FindById(ctx, conn, TestPilatesClass.Id)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestClassFindByName(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	m := Class{}
	err := m.FindByName(ctx, conn, TestPilatesClass.Name)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, TestPilatesClass.Name, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)
}

func TestValidateAPIClasses(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	c := 10
	notValidTime := time.Now().AddDate(-3, 0, 0)
	a := containers.APIClasses{
		Name:      "name",
		Capacity:  &c,
		StartDate: &notValidTime,
		EndDate:   &notValidTime,
	}
	class := Class{}
	err := class.ValidateAPIClasses(ctx, conn, a)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), constants.StartDateBeforeNowErrorMsg)

	now := time.Now().AddDate(0, 0, 1)
	future := time.Now().AddDate(3, 0, 0)
	a = containers.APIClasses{
		Name:      "name",
		Capacity:  &c,
		StartDate: &future,
		EndDate:   &now,
	}
	err = class.ValidateAPIClasses(ctx, conn, a)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), constants.StartDateAfterEndDateErrorMsg)

	date := time.Now().AddDate(0, 0, 10)
	id := uuid.New().String()
	v := TimetableItem{
		Id:      id,
		ClassId: TestYogaClass.Id,
		Date:    &date,
	}
	err = v.Insert(ctx, conn)
	assert.Nil(t, err)

	startDate := date
	endDate := date.AddDate(0, 0, 1)
	a = containers.APIClasses{
		Name:      "name",
		Capacity:  &c,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	err = class.ValidateAPIClasses(ctx, conn, a)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.TimetableItemDateExistsBatchErrorMsg, date.Format(constants.DateFormat)))

	err = v.Delete(ctx, conn)
	assert.Nil(t, err)

	err = class.ValidateAPIClasses(ctx, conn, a)
	assert.Nil(t, err)
}

func TestClassBatchInsert(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	startDate := time.Now().AddDate(0, 0, 2)
	endDate := time.Now().AddDate(0, 0, 21)
	c := 10
	name := "name"
	a := containers.APIClasses{
		Name:      name,
		Capacity:  &c,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	class := Class{}
	err := class.BatchInsert(ctx, conn, a)
	assert.Nil(t, err)
	ti := TimetableItem{}
	date := startDate
	counter := 0
	for !date.After(endDate) {
		err = ti.FindByDate(ctx, conn, &date)
		assert.Nil(t, err)
		plusDayDate := date.AddDate(0, 0, 1)
		date = plusDayDate
		counter++
	}
	assert.Equal(t, 20, counter)

	err = class.FindByName(ctx, conn, name)
	assert.Nil(t, err)
	assert.Equal(t, name, class.Name)
	assert.Equal(t, c, *class.Capacity)
}
