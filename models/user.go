package models

type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	SelfIncomeRefer int
	SelfIncome      SelfIncome `gorm:"foreignKey:SelfIncomeRefer"`

	Balance Balance
}
