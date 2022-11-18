package utils

import (
	"microservice-for-working-with-user-balance/config"
	"microservice-for-working-with-user-balance/models"
)

func GetUser(userID uint) models.User {
	var user models.User
	config.DB.Where("id = ?", userID).First(&user)
	config.DB.Preload("SelfIncomes").Find(&user)
	config.DB.Preload("Orders").Find(&user)
	config.DB.Preload("UsersTransfer").Find(&user)
	return user
}
