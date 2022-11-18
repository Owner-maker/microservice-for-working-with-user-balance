package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"microservice-for-working-with-user-balance/controllers"
)

func CreateEndpoints(route *gin.Engine) {
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
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
