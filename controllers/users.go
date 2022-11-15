package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userBalanceServicegot/models"
)

// GetUsers GET /users
func GetUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	context.JSON(http.StatusOK, gin.H{"users": users})
}
