package models

import (
	"fmt"
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
