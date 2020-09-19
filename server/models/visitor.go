package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
)

type Visitor struct {
	Id    string `json:"id" column:"id" validate:"required,uuid4" skip_on_update:"true"`
	Email string `json:"email" column:"email" validate:"required,email"`
}

func (Visitor) GetTableName() string {
	return constants.VisitorTableName
}

func (m Visitor) GetId() string {
	return m.Id
}

func (m *Visitor) FindById(db *sql.DB, id string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Email)
	return err
}
