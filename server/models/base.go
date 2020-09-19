package models

import (
	"database/sql"
	"github.com/vavilen84/class_booking/database"
	"github.com/vavilen84/class_booking/interfaces"
)

func Insert(db *sql.DB, s interface{}) (err error) {
	err = Validate(s)
	err = database.Insert(db, s.(interfaces.Model))
	return
}
