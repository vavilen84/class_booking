package models

import (
	"github.com/vavilen84/class_booking/helpers"
	"gopkg.in/go-playground/validator.v9"
)

func ValidateStruct(s interface{}) error {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		helpers.LogError(err)
		return err
	}
	return nil
}
