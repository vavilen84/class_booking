package models

import (
	"context"
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

func (m Visitor) Insert(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	existingVisitor := Visitor{}
	err = existingVisitor.FindByEmail(ctx, conn, m.Email)
	if err != sql.ErrNoRows {
		return errors.New(fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName))
	}
	err = database.Insert(ctx, conn, m)
	return
}

func (m Visitor) Update(ctx context.Context, conn *sql.Conn) (err error) {
	err = Validate(m)
	if err != nil {
		return
	}
	existingVisitor := Visitor{}
	err = existingVisitor.FindByEmail(ctx, conn, m.Email)
	if (err != sql.ErrNoRows) && (m.Id != existingVisitor.Id) {
		return errors.New(fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName))
	}
	err = database.Update(ctx, conn, m)
	return
}

func (m Visitor) Delete(ctx context.Context, conn *sql.Conn) {
	database.DeleteById(ctx, conn, m)
}

func (m *Visitor) FindById(ctx context.Context, conn *sql.Conn, id string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE id = ?`, id)
	err = row.Scan(&m.Id, &m.Email)
	return err
}

func (m *Visitor) FindByEmail(ctx context.Context, conn *sql.Conn, email string) (err error) {
	row := conn.QueryRowContext(ctx, `SELECT * FROM `+m.GetTableName()+` WHERE email = ?`, email)
	err = row.Scan(&m.Id, &m.Email)
	return err
}
