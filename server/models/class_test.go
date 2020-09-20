package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/store"
	"testing"
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
