package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Nickname string `json:"nickname" gorm:"unique;"`

	SelfIncomes   []SelfIncome    `gorm:"foreignKey:UserID; references:ID;"`
	UsersTransfer []UsersTransfer `gorm:"foreignKey:UserID; references:ID;"`
	Orders        []Order         `gorm:"foreignKey:UserID; references:ID;"`
	Balance       Balance
}
