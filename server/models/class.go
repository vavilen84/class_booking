package models

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
	"github.com/vavilen84/class_booking/helpers"
)

type Class struct {
	Id       string `json:"id" column:"id" validate:"required"`
	Name     string `json:"name" column:"name" validate:"required"`
	Capacity *int   `json:"capacity" column:"capacity" validate:"required"`
}

func (Class) GetTableName() string {
	return constants.ClassTableName
}

func (c *Class) Insert(db *sql.DB) (err error) {
	err = ValidateStruct(c)
	if err != nil {
		helpers.LogError(err)
		return
	}
	c.Id = uuid.New().String()
	err = database.Insert(db, c)
	return
}
