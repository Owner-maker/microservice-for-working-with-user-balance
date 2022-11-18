package config

import (
	"fmt"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "45129ff",
		DBName:   "balanceServiceDB",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}

func ConnectDB() {
	DBConfig := BuildDBConfig()
	db, err := gorm.Open("postgres", DbURL(DBConfig))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Balance{})
	db.AutoMigrate(&models.SelfIncome{})
	db.AutoMigrate(&models.UsersTransfer{})
	db.AutoMigrate(&models.Order{})

	DB = db
}
