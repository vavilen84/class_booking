package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/interfaces"
	"reflect"
	"strings"
)

func TxInsert(ctx context.Context, tx *sql.Tx, v interfaces.Model) error {

	reflectTypeOf := reflect.TypeOf(v)
	reflectValueOf := reflect.ValueOf(v)
	fieldsCount := reflectTypeOf.NumField()

	placeholders := make([]string, 0)
	values := make([]interface{}, 0)
	columns := make([]string, 0)
	paramCounter := 1

	// collect params, columns, values
	for i := 0; i < fieldsCount; i++ {
		field := reflectTypeOf.Field(i)
		skipOnInsert := field.Tag.Get("skip_on_insert")
		if skipOnInsert == "true" {
			continue
		}
		columns = append(columns, field.Tag.Get("column"))
		placeholders = append(placeholders, "?")
		values = append(values, reflectValueOf.FieldByName(field.Name).Interface())
		paramCounter++
	}

	// quote strings in values
	for i := 0; i < len(values); i++ {
		t := reflect.TypeOf(&values[i])
		v := reflect.ValueOf(&values[i])
		switch t.Kind() {
		case reflect.String:
			v.SetString("'" + v.String() + "'")
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		v.GetTableName(),
		strings.Join(columns, ","),
		strings.Join(placeholders, ","),
	)

	_, err := tx.ExecContext(ctx, query, values...)
	if err != nil {
		helpers.LogError(err)
		return err
	}
	return nil
}

func Insert(ctx context.Context, conn *sql.Conn, v interfaces.Model) error {

	reflectTypeOf := reflect.TypeOf(v)
	reflectValueOf := reflect.ValueOf(v)
	fieldsCount := reflectTypeOf.NumField()

	placeholders := make([]string, 0)
	values := make([]interface{}, 0)
	columns := make([]string, 0)
	paramCounter := 1

	// collect params, columns, values
	for i := 0; i < fieldsCount; i++ {
		field := reflectTypeOf.Field(i)
		skipOnInsert := field.Tag.Get("skip_on_insert")
		if skipOnInsert == "true" {
			continue
		}
		columns = append(columns, field.Tag.Get("column"))
		placeholders = append(placeholders, "?")
		values = append(values, reflectValueOf.FieldByName(field.Name).Interface())
		paramCounter++
	}

	// quote strings in values
	for i := 0; i < len(values); i++ {
		t := reflect.TypeOf(&values[i])
		v := reflect.ValueOf(&values[i])
		switch t.Kind() {
		case reflect.String:
			v.SetString("'" + v.String() + "'")
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		v.GetTableName(),
		strings.Join(columns, ","),
		strings.Join(placeholders, ","),
	)

	_, err := conn.ExecContext(ctx, query, values...)
	if err != nil {
		helpers.LogError(err)
		return err
	}
	return nil
}

func Update(ctx context.Context, conn *sql.Conn, v interfaces.Model) error {

	reflectTypeOf := reflect.TypeOf(v)
	reflectValueOf := reflect.ValueOf(v)
	fieldsCount := reflectTypeOf.NumField()

	values := make([]interface{}, 0)
	columns := make([]string, 0)
	paramCounter := 1

	// collect params, columns, values
	for i := 0; i < fieldsCount; i++ {
		field := reflectTypeOf.Field(i)
		skipOnUpdate := field.Tag.Get("skip_on_update")
		if skipOnUpdate == "true" {
			continue
		}
		columns = append(columns, field.Tag.Get("column"))
		values = append(values, reflectValueOf.FieldByName(field.Name).Interface())
		paramCounter++
	}

	// quote strings in values
	for i := 0; i < len(values); i++ {
		t := reflect.TypeOf(&values[i])
		v := reflect.ValueOf(&values[i])
		switch t.Kind() {
		case reflect.String:
			v.SetString("'" + v.String() + "'")
		}
	}

	var setData []string
	for _, v := range columns {
		setData = append(setData, v+" = ?")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = ?",
		v.GetTableName(),
		strings.Join(setData, ","),
	)
	values = append(values, v.GetId())
	_, err := conn.ExecContext(ctx, query, values...)
	if err != nil {
		helpers.LogError(err)
		return err
	}
	return nil
}

func DeleteById(ctx context.Context, conn *sql.Conn, v interfaces.Model) error {
	_, err := conn.ExecContext(ctx, `DELETE FROM `+v.GetTableName()+` WHERE id = ?`, v.GetId())
	return err
}
