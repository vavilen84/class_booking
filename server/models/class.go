package models

import (
	"github.com/vavilen84/class_booking/constants"
)

type Class struct {
	Id       string `json:"id" column:"id" validate:"required,uuid4"`
	Name     string `json:"name" column:"name" validate:"required,min=2,max=255"`
	Capacity *int   `json:"capacity" column:"capacity" validate:"required,numeric,min=1,max=50"`
}

func (Class) GetTableName() string {
	return constants.ClassTableName
}
