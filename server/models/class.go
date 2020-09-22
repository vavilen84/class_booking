package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
	"github.com/vavilen84/class_booking/helpers"
	"time"
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

func (m Class) Insert(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = database.Insert(ctx, conn, m)
	return
}

func (m Class) Update(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	err = database.Update(ctx, conn, m)
	return
}

func (m Class) Delete(ctx context.Context, conn *sql.Conn) error {
	return database.DeleteById(ctx, conn, m)
}

func (m *Class) FindById(ctx context.Context, conn *sql.Conn, id string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Name, &m.Capacity)
	return err
}

func (m *Class) FindByName(ctx context.Context, conn *sql.Conn, name string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE name LIKE ?`, name)
	err = row.Scan(&m.Id, &m.Name, &m.Capacity)
	return err
}

func (m Class) ValidateAPIClasses(ctx context.Context, conn *sql.Conn, apiClasses APIClasses) (err error) {
	err = Validate(apiClasses)
	if err != nil {
		return
	}
	if apiClasses.StartDate.After(*apiClasses.EndDate) {
		return errors.New(constants.StartDateAfterEndDateErrorMsg)
	}
	t := TimetableItem{}
	date := apiClasses.StartDate
	for !date.After(*apiClasses.EndDate) {
		err = t.FindByDate(ctx, conn, date)
		if err == nil {
			return errors.New(fmt.Sprintf(constants.TimetableItemDateExistsBatchErrorMsg, date.Format(constants.DateFormat)))
		}
		if err != sql.ErrNoRows {
			helpers.LogError(err)
			return
		}
		plusDayDate := date.AddDate(0, 0, 1)
		date = &plusDayDate
	}
	return nil
}

func (m Class) BatchInsert(ctx context.Context, conn *sql.Conn, apiClasses APIClasses) (err error) {
	err = m.ValidateAPIClasses(ctx, conn, apiClasses)
	if err != nil {
		helpers.LogError(err)
		return
	}
	tx, beginTxErr := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if beginTxErr != nil {
		helpers.LogError(beginTxErr)
		return beginTxErr
	}
	class := Class{}
	err = class.FindByName(ctx, conn, apiClasses.Name)
	if err == sql.ErrNoRows {
		class = Class{
			Id:       uuid.New().String(),
			Name:     apiClasses.Name,
			Capacity: apiClasses.Capacity,
		}
		err = Validate(class)
		if err != nil {
			_ = tx.Rollback()
			helpers.LogError(err)
			return
		}
		execErr := database.TxInsert(ctx, tx, class)
		if execErr != nil {
			_ = tx.Rollback()
			helpers.LogError(execErr)
			return execErr
		}
	}
	date := apiClasses.StartDate
	t := TimetableItem{}
	for !date.After(*apiClasses.EndDate) {
		err = t.FindByDate(ctx, conn, date)
		if err != sql.ErrNoRows {
			_ = tx.Rollback()
			err = errors.New(fmt.Sprintf(constants.TimetableItemDateExistsBatchErrorMsg, date.Format(constants.DateFormat)))
			helpers.LogError(err)
			return
		}
		t = TimetableItem{
			Id:      uuid.New().String(),
			Date:    date,
			ClassId: class.Id,
		}
		err = t.ValidateTimetableItemBeforeInsert(ctx, conn)
		if err != nil {
			_ = tx.Rollback()
			helpers.LogError(err)
			return
		}
		execErr := database.TxInsert(ctx, tx, t)
		if execErr != nil {
			_ = tx.Rollback()
			helpers.LogError(execErr)
			return execErr
		}
		plusDay := date.Add(24 * time.Hour)
		date = &plusDay
	}
	if err := tx.Commit(); err != nil {
		helpers.LogError(err)
		return err
	}
	return
}
