package main

import (
	"github.com/gin-gonic/gin"
	"microservice-for-working-with-user-balance/config"
	"microservice-for-working-with-user-balance/routes"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Balance Managing service
// @version 1.0
// @description This service is a task for the Avito internship. Provides you a REST API to work with user balances (crediting funds, debiting funds, transferring funds from user to user, as well as a method for obtaining a user's balance).

// @host localhost:8080
// @basePath /

func main() {
	route := gin.Default()
	config.ConnectDB()
	config.FillDbWithData()
	routes.CreateEndpoints(route)
	err := route.Run()
	if err != nil {
		return
	}
}
