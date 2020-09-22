package models

import (
	"context"
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

func (m TimetableItem) Insert(ctx context.Context, conn *sql.Conn) (err error) {
	err = m.ValidateTimetableItemBeforeInsert(ctx, conn)
	if err != nil {
		return
	}
	err = database.Insert(ctx, conn, m)
	return
}

func (m TimetableItem) ValidateTimetableItemBeforeInsert(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = m.ValidateDate(ctx, conn)
	if err != nil {
		return
	}
	err = m.ValidateClassExists(ctx, conn)
	if err != nil {
		return
	}
	return nil
}

func (m TimetableItem) Delete(ctx context.Context, conn *sql.Conn) error {
	return database.DeleteById(ctx, conn, m)
}

func (m *TimetableItem) FindById(ctx context.Context, conn *sql.Conn, id string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.ClassId, &m.Date)
	return err
}

func (m *TimetableItem) FindByDate(ctx context.Context, conn *sql.Conn, date *time.Time) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE date = ?`, date.Format(constants.DateFormat))
	err = row.Scan(&m.Id, &m.ClassId, &m.Date)
	return err
}

func (m TimetableItem) ValidateDate(ctx context.Context, conn *sql.Conn) error {
	existingClass := TimetableItem{}
	err := existingClass.FindByDate(ctx, conn, m.Date)
	if err != sql.ErrNoRows {
		return errors.New(fmt.Sprintf(constants.TimetableItemDateExistsErrorMsg, constants.TimetableItemStructName))
	}
	return nil
}

func (m TimetableItem) ValidateClassExists(ctx context.Context, conn *sql.Conn) error {
	class := Class{}
	err := class.FindById(ctx, conn, m.ClassId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(fmt.Sprintf(constants.ClassDoesNotExistErrorMsg, constants.TimetableItemStructName))
		}
		return err
	}
	return nil
}
