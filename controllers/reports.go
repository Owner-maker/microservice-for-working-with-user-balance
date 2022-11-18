package controllers

import (
	"fmt"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/utils"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"net/http"
	"os"
)

func UpdateServicesReport(context *gin.Context) {
	var input utils.GetServicesInfoInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := utils.GetServicesInfoFromDB(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to load information about services"})
		return
	}
	_, err = utils.CreateServicesInfoScv(*info)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to create csv file"})
		return
	}
	fmt.Println(context.Request.Host)
	context.JSON(http.StatusOK, gin.H{"report": fmt.Sprintf("http://%s/static/services", context.Request.Host)})
}

func GetServicesReport(context *gin.Context) {
	var serviceInfo []utils.ServiceInfo
	servicesFile, err := os.OpenFile("./static/services.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to provide report"})
		return
	}
	defer servicesFile.Close()

	err = gocsv.UnmarshalFile(servicesFile, &serviceInfo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to provide report"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"report": serviceInfo})
}
