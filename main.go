package main

import (
	"github.com/gin-gonic/gin"
	"userBalanceServicegot/controllers"
	"userBalanceServicegot/models"
)

func main() {
	route := gin.Default()

	models.ConnectDB()
	models.FillDbWithData()

	route.GET("/users", controllers.GetUsers)
	route.GET("/user/balance", controllers.GetUserBalance)
	route.PATCH("/user/balance/topup", controllers.UpdateUserBalance)
	route.PATCH("/users/transfer", controllers.AccomplishUsersTransfer)
	route.POST("/user/buy/service", controllers.CreateOrder)
	route.PATCH("/user/perform/service", controllers.PerformService)

	err := route.Run()
	if err != nil {
		return
	}
}
