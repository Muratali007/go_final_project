package models

type Footballer struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Price int64 `json:"price"`
	Club string `json:"club"`
}
