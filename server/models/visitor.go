package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
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

func (m Visitor) Insert(db *sql.DB) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	existingVisitor := Visitor{}
	err = existingVisitor.FindByEmail(db, m.Email)
	if err != sql.ErrNoRows {
		return errors.New(fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName))
	}
	err = database.Insert(db, m)
	return
}

func (m Visitor) Update(db *sql.DB) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	existingVisitor := Visitor{}
	err = existingVisitor.FindByEmail(db, m.Email)
	if (err != sql.ErrNoRows) && (m.Id != existingVisitor.Id) {
		return errors.New(fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName))
	}
	err = database.Update(db, m)
	return
}

func (m Visitor) Delete(db *sql.DB) {
	database.DeleteById(db, m)
}

func (m *Visitor) FindById(db *sql.DB, id string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Email)
	return err
}

func (m *Visitor) FindByEmail(db *sql.DB, email string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE email = ?`, email)
	err = row.Scan(&m.Id, &m.Email)
	return err
}
