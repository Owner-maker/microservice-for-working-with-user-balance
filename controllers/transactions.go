package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userBalanceServicegot/config"
	"userBalanceServicegot/utils"
)

type UserTransactionsInput struct {
	UserID uint   `json:"user_id"`
	Limit  uint   `json:"limit"`
	Page   uint   `json:"page"`
	Sort   string `json:"sort"`
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
