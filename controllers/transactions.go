package controllers

import (
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserTransactionsInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Limit  uint   `json:"limit" binding:"required"`
	Page   uint   `json:"page" binding:"required"`
	Sort   string `json:"sort" binding:"required"`
}

func GetPaginatedUsersTransactions(context *gin.Context) {
	var input UserTransactionsInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserID).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	pagination := utils.Pagination{
		Limit: input.Limit,
		Page:  input.Page,
		Sort:  input.Sort,
	}

	transactions, err := utils.GetPaginatedUserTransactions(input.UserID, pagination)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
