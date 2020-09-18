package models

type Class struct {
	Id       string `json:"id" column:"created_at" validate:"required"`
	Name     string `json:"name" column:"created_at" validate:"required"`
	Date     string `json:"date" column:"created_at" validate:"required"`
	Capacity int    `json:"capacity" column:"created_at" validate:"required"`
}
