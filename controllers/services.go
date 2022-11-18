package controllers

import (
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ReserveMoneyForServiceInput struct {
	UserID    uint `json:"user_id" binding:"required"`
	ServiceID uint `json:"service_id" binding:"required"`
	Price     uint `json:"price" binding:"required"`
}

type HandleServiceInput struct {
	UserID    uint `json:"user_id" binding:"required"`
	ServiceID uint `json:"service_id" binding:"required"`
	OrderID   uint `json:"order_id" binding:"required"`
}

func CreateOrder(context *gin.Context) {
	var input ReserveMoneyForServiceInput
	var balance models.Balance
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserID).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	config.DB.Where("user_id = ?", input.UserID).First(&balance)
	if balance.Value-input.Price < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User does not have enough money to buy the service"})
		return
	}
	order := models.Order{
		UserID:          input.UserID,
		ServiceID:       input.ServiceID,
		IncomingBalance: balance.Value,
		OutgoingBalance: balance.Value - input.Price,
		Price:           input.Price,
	}

	balance.Value -= input.Price
	config.DB.Model(&balance).Update(&balance)
	config.DB.Create(&order)
	context.JSON(http.StatusOK, gin.H{"order_id": order.ID, "balance": balance.Value})
}

func PerformService(context *gin.Context) {
	var input HandleServiceInput
	var order models.Order
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserID).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	if err := config.DB.Where("id = ?", input.OrderID).First(&order).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such order"})
		return
	}
	if input.ServiceID != order.ServiceID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such service id with this order id"})
		return
	}
	if order.IsCompleted == true {
		context.JSON(http.StatusBadRequest, gin.H{"error": "order was already completed"})
		return
	}
	order.Timestamp = time.Now()
	order.IsCompleted = true
	config.DB.Model(&order).Update(&order)
	context.Status(200)
}
func CancelService(context *gin.Context) {
	var input HandleServiceInput
	var order models.Order
	var balance models.Balance
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserID).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	if err := config.DB.Where("id = ?", input.OrderID).First(&order).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such order"})
		return
	}
	if input.ServiceID != order.ServiceID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such service id with this order id"})
		return
	}
	config.DB.Where("user_id = ?", input.UserID).First(&balance)
	order.IncomingBalance = balance.Value
	order.OutgoingBalance = order.IncomingBalance + order.Price
	order.Timestamp = time.Now()
	balance.Value += order.Price
	config.DB.Model(&balance).Update(&balance)
	config.DB.Model(&order).Update(&order)
	context.Status(200)
}
