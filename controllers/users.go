package controllers

import (
	"fmt"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"

	_ "github.com/Owner-maker/microservice-for-working-with-user-balance/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type ErrorOutput struct {
	Error string `json:"error"`
}

type BalanceInfoOutput struct {
	Balance uint `json:"balance"`
}

type GetUserBalanceInput struct {
	ID uint `json:"id" binding:"required"`
}

type GetUserInput struct {
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

func GetUser(context *gin.Context) {
	var user models.User
	var input GetUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	config.DB.Preload("SelfIncomes").Find(&user)
	config.DB.Preload("Orders").Find(&user)
	config.DB.Preload("UsersTransfer").Find(&user)
	context.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary GetUserBalance
// @Description Method allows you to get user's balance value via id
// @ID get-users-balance
// @Tags users
// @Accept json
// @Produce json
// @Param input body GetUserBalanceInput true "User's balance info"
// @Success 200 {object} BalanceInfoOutput
// @Failure 400 {object} ErrorOutput
// @Router /user/balance [post]
func GetUserBalance(context *gin.Context) {
	var user models.User
	var balance models.Balance
	var input GetUserBalanceInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	if err := config.DB.Where("user_id = ?", input.ID).First(&balance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user has not a balance yet"})
		return
	}
	context.JSON(http.StatusOK, BalanceInfoOutput{Balance: balance.Value})
}

// @Summary UpdateUserBalance
// @Description Method allows you to top up user's balance value via id and create transaction
// @ID update-user-balance
// @Tags users
// @Accept json
// @Produce json
// @Param input body UpdateUserBalanceInput true "Info to top up user's balance"
// @Success 200 {object} BalanceInfoOutput
// @Failure 400 {object} ErrorOutput
// @Router /user/balance/topup [patch]
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
	if err := config.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err = config.DB.Where("user_id = ?", input.ID).First(&balance).Error; err != nil {
		newBalance := models.Balance{UserID: input.ID, Value: 0}
		config.DB.Create(&newBalance)
	}
	CreateSelfIncomeTransaction(input)
	config.DB.Where("user_id = ?", input.ID).First(&balance)
	balance.Value += uint(input.Value)
	config.DB.Model(&balance).Update(&balance)
	context.JSON(http.StatusOK, BalanceInfoOutput{Balance: balance.Value})
}

// @Summary AccomplishUsersTransfer
// @Description Method allows you to sen money to another user
// @ID accomplish-users-transfer
// @Tags users
// @Accept json
// @Produce json
// @Param input body UserTransferInput true "Info to send money to user"
// @Success 200 {object} BalanceInfoOutput
// @Failure 400 {object} ErrorOutput
// @Router /users/transfer [patch]
func AccomplishUsersTransfer(context *gin.Context) {
	var input UserTransferInput
	var userSender models.User
	var userGetter models.User
	var userGetterBalance models.Balance
	var userSenderBalance models.Balance
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserSenderID).First(&userSender).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("there is no such user with id = %d", input.UserSenderID)})
		return
	}
	if err := config.DB.Where("id = ?", input.UserGetterID).First(&userGetter).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("there is no such user with id = %d", input.UserGetterID)})
		return
	}
	if err := config.DB.Where("user_id = ?", userSender.ID).First(&userSenderBalance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user with id = %d has not a balance yet", userSender.ID)})
		return
	}
	if err := config.DB.Where("user_id = ?", userGetter.ID).First(&userGetterBalance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user with id = %d has not a balance yet", userGetter.ID)})
		return
	}
	if int(userSenderBalance.Value)-int(input.Value) < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User sender does not have enough money to make transfer"})
		return
	}
	CreateUserTransferTransaction(userSenderBalance, userGetter.ID, -int(input.Value))
	CreateUserTransferTransaction(userGetterBalance, userSender.ID, int(input.Value))

	userSenderBalance.Value -= input.Value
	userGetterBalance.Value += input.Value

	config.DB.Model(&userSenderBalance).Update(&userSenderBalance)
	config.DB.Model(&userGetterBalance).Update(&userGetterBalance)
	context.JSON(http.StatusOK, BalanceInfoOutput{Balance: userSenderBalance.Value})
}

func CreateSelfIncomeTransaction(updateInput UpdateUserBalanceInput) {
	var currentBalance models.Balance
	config.DB.Where("user_id = ?", updateInput.ID).First(&currentBalance)

	income := models.SelfIncome{
		UserID:          updateInput.ID,
		IncomingBalance: currentBalance.Value,
		OutgoingBalance: currentBalance.Value + uint(updateInput.Value),
		Timestamp:       time.Now(),
		MoneyValue:      uint(updateInput.Value),
	}

	config.DB.Create(&income)
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

	config.DB.Create(&transfer)
}
