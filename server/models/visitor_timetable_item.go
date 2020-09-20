package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
)

type VisitorTimetableItem struct {
	Id              string `json:"id" column:"id" validate:"required,uuid4" skip_on_update:"true"`
	VisitorId       string `json:"visitor_id" column:"visitor_id" validate:"required,uuid4"`
	TimetableItemId string `json:"timetable_item_id" column:"timetable_item_id" validate:"required,uuid4"`
}

func (VisitorTimetableItem) GetTableName() string {
	return constants.VisitorTimetableItemTableName
}

func (m VisitorTimetableItem) GetId() string {
	return m.Id
}

func (m VisitorTimetableItem) Insert(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = m.ValidateVisitorExists(ctx, conn)
	if err != nil {
		return
	}
	err = m.ValidateTimetableItemExists(ctx, conn)
	if err != nil {
		return
	}
	err, alreadyExists := m.BookingByVisitorAndTimetableItemExists(ctx, conn)
	if err != nil {
		return
	}
	if alreadyExists {
		return errors.New(fmt.Sprintf(constants.BookingAlreadyExistsErrorMsg, constants.VisitorTimetableItemStructName))
	}
	err = database.Insert(ctx, conn, m)
	return
}

func (m VisitorTimetableItem) Delete(ctx context.Context, conn *sql.Conn) {
	database.DeleteById(ctx, conn, m)
}

func (m *VisitorTimetableItem) FindById(ctx context.Context, conn *sql.Conn, id string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.VisitorId, &m.TimetableItemId)
	return err
}

func (m *VisitorTimetableItem) BookingByVisitorAndTimetableItemExists(ctx context.Context, conn *sql.Conn) (err error, exists bool) {
	row := conn.QueryRowContext(
		ctx,
		`SELECT * FROM `+m.GetTableName()+` WHERE visitor_id = ? AND timetable_item_id = ?`,
		m.VisitorId,
		m.TimetableItemId,
	)
	err = row.Scan(&m.Id, &m.VisitorId, &m.TimetableItemId)
	if err == nil {
		return nil, true
	}
	if err == sql.ErrNoRows {
		return nil, false
	}
	return
}

func (m *VisitorTimetableItem) ValidateVisitorExists(ctx context.Context, conn *sql.Conn) error {
	v := Visitor{}
	err := v.FindById(ctx, conn, m.VisitorId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(fmt.Sprintf(constants.VisitorDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
		}
		return err
	}
	return nil
}

func (m *VisitorTimetableItem) ValidateTimetableItemExists(ctx context.Context, conn *sql.Conn) error {
	v := TimetableItem{}
	err := v.FindById(ctx, conn, m.TimetableItemId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(fmt.Sprintf(constants.TimetableItemDoesNotExistErrorMsg, constants.VisitorTimetableItemStructName))
		}
		return err
	}
	return nil
}
