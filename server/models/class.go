package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
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

func (m *Class) FindById(db *sql.DB, id string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Name, &m.Capacity)
	return err
}
