package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
	"userBalanceServicegot/models"
)

type GetUserBalanceInput struct {
	ID uint `json:"id" binding:"required"`
}

type UpdateUserBalanceInput struct {
	ID    uint `json:"id" binding:"required"`
	Value int  `json:"value" binding:"required"`
}

type UserTransferInput struct {
	UserSenderID uint `json:"user_sender_id" binding:"required"`
	UserGetterID uint `json:"user_getter_id" binding:"required"`
	Value        uint `json:"value" binding:"required"`
}

func GetUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	//models.DB.Preload("SelfIncomes").Find(&users)
	//models.DB.Preload("Balance").Find(&users)
	context.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserBalance(context *gin.Context) {
	var user models.User
	var balance models.Balance
	var input GetUserBalanceInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	if err := models.DB.Where("user_id = ?", input.ID).First(&balance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user has not a balance yet"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"balance": balance.Value})
}

func UpdateUserBalance(context *gin.Context) {
	var input UpdateUserBalanceInput
	var balance models.Balance
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Value <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "money value can not be zero or negative"})
		return
	}
	user := models.User{ID: input.ID}
	if err := models.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		newBalance := models.Balance{UserID: input.ID, Value: 0}
		models.DB.Create(&newBalance)
	}
	CreateSelfIncomeTransaction(input)
	models.DB.Where("user_id = ?", input.ID).First(&balance)
	balance.Value += uint(input.Value)
	models.DB.Model(&balance).Update(&balance)
	context.JSON(http.StatusOK, gin.H{"balance": balance})
}

func AccomplishUsersTransfer(context *gin.Context) {
	var input UserTransferInput
	var userSender models.User
	var userGetter models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("id = ?", input.UserSenderID).First(&userSender).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("id = ?", input.UserGetterID).First(&userGetter).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Preload("Balance").Find(&userSender)
	if userSender.Balance.Value-input.Value < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User sender does not have enough money to make transfer"})
		return
	}
	models.DB.Preload("Balance").Find(&userGetter)
	userSender.Balance.Value -= input.Value
	userGetter.Balance.Value += input.Value
	models.DB.Model(&userSender.Balance).Update(&userSender.Balance)
	models.DB.Model(&userGetter.Balance).Update(&userGetter.Balance)

	CreateUserTransferTransaction(userSender.Balance, userGetter.ID, -int(input.Value))
	CreateUserTransferTransaction(userGetter.Balance, userSender.ID, int(input.Value))
	context.JSON(http.StatusOK, gin.H{"balance_user_sender": userSender.Balance.Value})
}

func CreateSelfIncomeTransaction(updateInput UpdateUserBalanceInput) {
	var currentBalance models.Balance
	models.DB.Where("user_id = ?", updateInput.ID).First(&currentBalance)

	income := models.SelfIncome{
		UserID:          updateInput.ID,
		IncomingBalance: currentBalance.Value,
		OutgoingBalance: currentBalance.Value + uint(updateInput.Value),
		Timestamp:       time.Now(),
		MoneyValue:      uint(updateInput.Value),
	}

	models.DB.Create(&income)
}

func CreateUserTransferTransaction(userBalance models.Balance, anotherUserID uint, value int) {
	transfer := models.UsersTransfer{
		UserID:          userBalance.UserID,
		AnotherUserID:   anotherUserID,
		IncomingBalance: userBalance.Value,
		OutgoingBalance: uint(int(userBalance.Value) + value),
		Timestamp:       time.Now(),
		MoneyValue:      uint(math.Abs(float64(value))),
	}

	models.DB.Create(&transfer)
}
