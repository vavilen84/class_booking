package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/store"
	"testing"
)

func TestVisitorValidateRequiredTag(t *testing.T) {
	err := Validate(Visitor{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorStructName, "Id"))
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.RequiredErrorMsg, constants.VisitorStructName, "Email"))
}

func TestClassValidateEmailTag(t *testing.T) {
	v := Visitor{
		Email: "not_valid_email",
	}
	err := Validate(v)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.EmailErrorMsg, constants.VisitorStructName))
}

func TestVisitorValidateUuid4Tag(t *testing.T) {
	v := Visitor{
		Id: "not valid uuid4",
	}
	err := Validate(v)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf(constants.Uuid4ErrorMsg, constants.VisitorStructName, "Id"))
}

func TestVisitorInsert(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	email := "new_visitor@example.com"
	v := Visitor{
		Id:    id,
		Email: email,
	}
	err := v.Insert(ctx, conn)
	assert.Nil(t, err)

	v = Visitor{}
	err = v.FindById(ctx, conn, id)
	assert.Nil(t, err)
	assert.Equal(t, v.Id, id)
	assert.Equal(t, v.Email, email)
}

func TestVisitorInsertWithAlreadyRegisteredEmail(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	v := Visitor{
		Id:    id,
		Email: TestVisitor.Email,
	}
	err := v.Insert(ctx, conn)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName), err.Error())
}

func TestVisitorUpdateWithAlreadyRegisteredEmail(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	id := uuid.New().String()
	email := "new_visitor@example.com"
	v := Visitor{
		Id:    id,
		Email: email,
	}
	err := v.Insert(ctx, conn)
	assert.Nil(t, err)

	v.Email = TestVisitor.Email
	err = v.Update(ctx, conn)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constants.VisitorEmailExistsErrorMsg, constants.VisitorStructName), err.Error())
}

func TestVisitorFindById(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	v := Visitor{}
	err := v.FindById(ctx, conn, TestVisitor.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitor.Id, v.Id)
	assert.Equal(t, TestVisitor.Email, v.Email)
}

func TestVisitorUpdate(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	v := Visitor{}
	err := v.FindById(ctx, conn, TestVisitor.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitor.Id, v.Id)
	assert.Equal(t, TestVisitor.Email, v.Email)

	updatedEmail := "updated_email@example.com"
	v.Email = updatedEmail
	err = v.Update(ctx, conn)
	assert.Nil(t, err)
	assert.Equal(t, updatedEmail, v.Email)

	v = Visitor{}
	err = v.FindById(ctx, conn, TestVisitor.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitor.Id, v.Id)
	assert.Equal(t, updatedEmail, v.Email)
}

func TestVisitorDelete(t *testing.T) {
	ctx := store.GetDefaultDBContext()
	conn := store.GetNewTestDBConn()
	defer conn.Close()
	PrepareTestDB(ctx, conn)

	v := Visitor{}
	err := v.FindById(ctx, conn, TestVisitor.Id)
	assert.Nil(t, err)
	assert.Equal(t, TestVisitor.Id, v.Id)
	assert.Equal(t, TestVisitor.Email, v.Email)

	err = v.Delete(ctx, conn)
	assert.Nil(t, err)

	err = v.FindById(ctx, conn, TestVisitor.Id)
	assert.Equal(t, sql.ErrNoRows, err)
}
