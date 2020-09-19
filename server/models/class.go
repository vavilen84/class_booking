package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
)

type Class struct {
	Id       string `json:"id" column:"id" validate:"required,uuid4" skip_on_update:"true"`
	Name     string `json:"name" column:"name" validate:"required,min=2,max=255"`
	Capacity *int   `json:"capacity" column:"capacity" validate:"required,numeric,min=1,max=50"`
}

func (Class) GetTableName() string {
	return constants.ClassTableName
}

func (m Class) GetId() string {
	return m.Id
}

func (m Class) Insert(db *sql.DB) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = database.Insert(db, m)
	return
}

func (m Class) Update(db *sql.DB) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = database.Update(db, m)
	return
}

func (m Class) Delete(db *sql.DB) {
	database.DeleteById(db, m)
}

func (m *Class) FindById(db *sql.DB, id string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Name, &m.Capacity)
	return err
}
