package main

import (
	"github.com/gin-gonic/gin"
	"userBalanceServicegot/controllers"
	"userBalanceServicegot/models"
)

func main() {
	route := gin.Default()

	models.ConnectDB()
	//models.FillDbWithData()

	route.GET("/users", controllers.GetUsers)
	route.GET("/user/balance", controllers.GetUserBalance)
	route.PATCH("/user/balance/topup", controllers.UpdateUserBalance)

	err := route.Run()
	if err != nil {
		return
	}
}
