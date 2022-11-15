package models

type Balance struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id"`
	Value  uint `json:"value"`
}
