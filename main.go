package main

import (
	"github.com/gin-gonic/gin"
	"userBalanceServicegot/config"
	"userBalanceServicegot/models"
	"userBalanceServicegot/routes"
)

func main() {
	route := gin.Default()
	config.ConnectDB()
	models.FillDbWithData()
	routes.CreateEndpoints(route)
	err := route.Run()
	if err != nil {
		return
	}
}
