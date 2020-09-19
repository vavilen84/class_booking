package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
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
