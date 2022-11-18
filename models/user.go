package models

type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	SelfIncomes   []SelfIncome    `gorm:"foreignKey:UserID; references:ID;"`
	UsersTransfer []UsersTransfer `gorm:"foreignKey:UserID; references:ID;"`
	Orders        []Order         `gorm:"foreignKey:UserID; references:ID;"`
	Balance       Balance
}
