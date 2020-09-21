package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/helpers"
	"testing"
)

func TestAPIClassesValidateRequiredTag(t *testing.T) {
	err := Validate(APIClasses{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "StartDate"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "Name"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "Capacity"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "EndDate"))
}

func TestAPIClassesValidateMinValueTag(t *testing.T) {
	notValidCapacity := 0
	c := APIClasses{
		Name:     "n",
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, constants.APIClassesStructName, "Name", "2"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MinValueErrorMsg, constants.APIClassesStructName, "Capacity", "1"))
}

func TestAPIClassesValidateMaxValueTag(t *testing.T) {
	notValidCapacity := 51
	c := APIClasses{
		Name:     helpers.GenerateRandomString(256),
		Capacity: &notValidCapacity,
	}
	err := Validate(c)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, constants.APIClassesStructName, "Name", "255"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.MaxValueErrorMsg, constants.APIClassesStructName, "Capacity", "50"))
}
