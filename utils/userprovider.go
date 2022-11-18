package utils

import (
	"userBalanceServicegot/config"
	"userBalanceServicegot/models"
)

func GetUser(userID uint) models.User {
	var user models.User
	config.DB.Where("id = ?", userID).First(&user)
	config.DB.Preload("SelfIncomes").Find(&user)
	config.DB.Preload("Orders").Find(&user)
	config.DB.Preload("UsersTransfer").Find(&user)
	return user
}
