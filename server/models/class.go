package models

type Class struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Capacity int    `json:"capacity"`
}
