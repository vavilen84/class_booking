package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"testing"
	"time"
)

func TestAPIBookingsValidateRequiredTag(t *testing.T) {
	err := Validate(APIBookings{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIBookingsStructName, "Date"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.APIBookingsStructName, "Email"))
}

func TestAPIBookingsValidateEmailTag(t *testing.T) {
	now := time.Now()
	v := APIBookings{
		Date:  &now,
		Email: "not_valid_email",
	}
	err := Validate(v)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.EmailErrorMsg, constants.APIBookingsStructName))
}
