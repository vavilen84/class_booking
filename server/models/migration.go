package models

import "github.com/vavilen84/class_booking/constants"

type Migration struct {
	Version   int64  `json:"version" column:"version" validate:"required"`
	Filename  string `json:"filename" column:"filename" validate:"required"`
	CreatedAt int64  `json:"created_at" column:"created_at" validate:"required"`
}

func (Migration) GetTableName() string {
	return constants.MigrationsTableName
}
