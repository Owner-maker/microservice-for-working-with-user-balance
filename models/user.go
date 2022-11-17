package models

type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	SelfIncomes   []*SelfIncome `gorm:"foreignKey:UserID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UsersTransfer []UsersTransfer
	Order         []Order
	Balance       Balance
}
