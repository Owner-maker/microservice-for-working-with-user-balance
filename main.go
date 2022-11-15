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

	err := route.Run()
	if err != nil {
		return
	}
}
