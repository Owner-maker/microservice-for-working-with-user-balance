package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=balanceServiceDB password=45129ff sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Balance{})
	db.AutoMigrate(&SelfIncome{})
	db.AutoMigrate(&UsersTransfer{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Service{})

	DB = db
}
