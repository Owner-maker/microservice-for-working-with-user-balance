package controllers

import (
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/models"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetPaginatedUsersTransactionsOutput struct {
	Transactions *[]utils.UserFormattedTransaction `json:"transactions"`
}

type UserTransactionsInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Limit  uint   `json:"limit" binding:"required"`
	Page   uint   `json:"page" binding:"required"`
	Sort   string `json:"sort" binding:"required"`
}

// @Summary GetPaginatedUsersTransactions
// @Description Method allows to get user's transactions info using the pagination, it allows to order transactions by date, money and other transaction's attributes. Limit - maximum of needed transactions, Page - the offset with limit, Sort - value to sort by:date: "timestamp asc" or money: "money_value asc" (also "desc" is available)
// @ID get-paginated-users-transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Param input body UserTransactionsInput true "Information to get transactions"
// @Success 200 {object} GetPaginatedUsersTransactionsOutput
// @Failure 400 {object} ErrorOutput
// @Router /user/transactions [post]
func GetPaginatedUsersTransactions(context *gin.Context) {
	var input UserTransactionsInput
	var user models.User
	var balance models.Balance
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "there is no such user"})
		return
	}
	if err := config.DB.Where("user_id = ?", input.UserID).First(&balance).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user has not a balance yet"})
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
	context.JSON(http.StatusOK, GetPaginatedUsersTransactionsOutput{Transactions: transactions})
}
