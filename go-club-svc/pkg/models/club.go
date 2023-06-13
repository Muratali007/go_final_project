package models

type Club struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Year int64 `json:"year"`
	Titles int64 `json:"titles"`
}