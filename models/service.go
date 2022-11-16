package models

type Service struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price uint   `json:"price"`

	Order []Order
}
