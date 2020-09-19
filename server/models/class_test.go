package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
	"testing"
)

func TestClassValidateRequiredTag(t *testing.T) {
	err := Validate(Class{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, "Class", "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, "Class", "Name"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, "Class", "Capacity"))
}

func TestClassValidateMinValueTag(t *testing.T) {
	notValidCapacity := 0
	c := Class{
		Name:     "n",
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, "Class", "Name", "2"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, "Class", "Capacity", "1"))
}

func TestClassValidateUuid4Tag(t *testing.T) {
	c := Class{
		Id: "not valid uuid4",
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, "Class"))
}

func TestClassValidateMaxValueTag(t *testing.T) {
	notValidCapacity := 51
	c := Class{
		Name:     generateRandomString(256),
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, "Class", "Name", "255"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, "Class", "Capacity", "50"))
}

func TestClassInsert(t *testing.T) {
	db := PrepareTestDB()
	capacity := 10
	id := uuid.New().String()
	c := Class{
		Id:       id,
		Name:     "Crossfit",
		Capacity: &capacity,
	}
	err := Insert(db, c)
	assert.Nil(t, err)
}

func TestClassFindByIdAndUpdate(t *testing.T) {
	db := PrepareTestDB()
	m := Class{}
	err := m.FindById(db, TestPilatesClass.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestPilatesClass.Id, m.Id)
	assert.Equal(t, TestPilatesClass.Name, m.Name)
	assert.Equal(t, TestPilatesClass.Capacity, m.Capacity)

	updatedName := "Updated Name"
	m.Name = updatedName
	err = database.Update(db, m)
	assert.Nil(t, err)
	assert.Equal(t, updatedName, m.Name)
}
