package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
	"time"
)

type TimetableItem struct {
	Id      string     `json:"id" column:"id" validate:"required,uuid4" skip_on_update:"true"`
	ClassId string     `json:"class_id" column:"class_id" validate:"required,uuid4"`
	Date    *time.Time `json:"date" column:"date" validate:"required"`
}

func (TimetableItem) GetTableName() string {
	return constants.TimetableItemTableName
}

func (m TimetableItem) GetId() string {
	return m.Id
}

func (m TimetableItem) Insert(db *sql.DB) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = m.ValidateDate(db)
	if err != nil {
		return
	}
	err = database.Insert(db, m)
	return
}

func (m TimetableItem) Delete(db *sql.DB) {
	database.DeleteById(db, m)
}

func (m *TimetableItem) FindById(db *sql.DB, id string) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.ClassId, &m.Date)
	return err
}

func (m *TimetableItem) FindByDate(db *sql.DB, date *time.Time) (err error) {
	row := db.QueryRow(`SELECT * FROM `+m.GetTableName()+` WHERE date = ?`, date.Format(constants.DateFormat))
	err = row.Scan(&m.Id, &m.ClassId, &m.Date)
	return err
}

func (m TimetableItem) ValidateDate(db *sql.DB) error {
	existingClass := TimetableItem{}
	err := existingClass.FindByDate(db, m.Date)
	if err != sql.ErrNoRows {
		return errors.New(fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
	}
	return nil
}
