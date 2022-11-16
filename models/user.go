package models

type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	SelfIncome    []SelfIncome
	UsersTransfer []UsersTransfer
	Order         []Order
	Balance       Balance
}
