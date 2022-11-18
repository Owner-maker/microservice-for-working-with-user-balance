package main

import (
	"github.com/gin-gonic/gin"
	"userBalanceServicegot/config"
	"userBalanceServicegot/controllers"
	"userBalanceServicegot/models"
)

func main() {
	route := gin.Default()

	config.ConnectDB()
	models.FillDbWithData()

	route.GET("/user", controllers.GetUser)
	route.GET("/user/balance", controllers.GetUserBalance)
	route.PATCH("/user/balance/topup", controllers.UpdateUserBalance)
	route.PATCH("/users/transfer", controllers.AccomplishUsersTransfer)
	route.POST("/user/buy/service", controllers.CreateOrder)
	route.PATCH("/user/perform/service", controllers.PerformService)
	route.PATCH("/user/cancel/service", controllers.CancelService)
	route.GET("/user/transactions", controllers.GetPaginatedUsersTransactions)
	route.GET("/services/report", controllers.UpdateServicesReport)
	route.GET("/static/services", controllers.GetServicesReport)

	err := route.Run()
	if err != nil {
		return
	}
}
