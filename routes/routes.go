package routes

import (
	"github.com/Owner-maker/microservice-for-working-with-user-balance/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateEndpoints(route *gin.Engine) {
	route.POST("/user", controllers.GetUser)
	route.POST("/user/balance", controllers.GetUserBalance)
	route.PATCH("/user/balance/topup", controllers.UpdateUserBalance)
	route.PATCH("/users/transfer", controllers.AccomplishUsersTransfer)
	route.POST("/user/buy/service", controllers.CreateOrder)
	route.PATCH("/user/perform/service", controllers.PerformService)
	route.DELETE("/user/cancel/service", controllers.CancelService)
	route.POST("/user/transactions", controllers.GetPaginatedUsersTransactions)
	route.POST("/services/report", controllers.UpdateServicesReport)
	route.POST("/static/services", controllers.GetServicesReport)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
