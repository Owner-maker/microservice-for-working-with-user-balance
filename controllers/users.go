package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"userBalanceServicegot/models"
)

type GetUserBalanceInput struct {
	ID string `json:"id" binding:"required"`
}

type UpdateUserBalanceInput struct {
	ID    string `json:"id" binding:"required"`
	Value uint   `json:"value" binding:"required"`
}

func GetUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	context.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserBalance(context *gin.Context) {
	var balance models.Balance
	var input GetUserBalanceInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("user_id = ?", input.ID).First(&balance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user has not a balance yet"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"balance": balance.Value})
}

func UpdateUserBalance(context *gin.Context) {
	var balance models.Balance
	var input UpdateUserBalanceInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("user_id = ?", input.ID).First(&balance).Error; err != nil {
		id, err := strconv.ParseInt(input.ID, 10, 32)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newBalance := models.Balance{UserID: uint(id), Value: 0}
		models.DB.Create(&newBalance)
	}
	balance.Value += input.Value
	models.DB.Model(&balance).Update(&balance)
	context.JSON(http.StatusOK, gin.H{"balance": balance})
}
