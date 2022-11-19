package controllers

import (
	"fmt"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/utils"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"net/http"
	"os"
)

type UpdateServicesReportOutput struct {
	Report string `json:"report"`
}

type GetServicesReportOutput struct {
	Report []utils.ServiceInfo `json:"report"`
}

// @Summary UpdateServicesReport
// @Description Method allows to generate new scv file on server of all sold services with its total sum via the information of Year and Month; returns link to reposrt's info (json)
// @ID update-service-sreport
// @Tags transactions
// @Accept json
// @Produce json
// @Param input body utils.GetServicesInfoInput true "Information to generate new scv report"
// @Success 200 {object} UpdateServicesReportOutput
// @Failure 400 {object} ErrorOutput
// @Router /services/report [post]
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
	context.JSON(http.StatusOK, UpdateServicesReportOutput{Report: fmt.Sprintf("http://%s/static/services", context.Request.Host)})
}

// @Summary GetServicesReport
// @Description Method allows to get information about all sold services from the generated scv file
// @ID get-services-report
// @Tags transactions
// @Accept json
// @Produce json
// @Success 200 {object} GetServicesReportOutput
// @Failure 400 {object} ErrorOutput
// @Router /static/services [post]
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
	context.JSON(http.StatusOK, GetServicesReportOutput{Report: serviceInfo})
}
