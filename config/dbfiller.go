package config

import (
	"userBalanceServicegot/models"
)

func FillDbWithData() {
	user0 := models.User{Nickname: "John Malkovich"}
	user1 := models.User{Nickname: "Silvester Stallone"}
	DB.Create(&user0)
	DB.Create(&user1)
}
