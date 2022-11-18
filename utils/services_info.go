package utils

import (
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"github.com/gocarina/gocsv"
	"os"
	"strconv"
)

type GetServicesInfoInput struct {
	Year  uint `json:"year" binding:"required"`
	Month uint `json:"month" binding:"required"`
}

type ServiceInfo struct {
	ServiceID uint `csv:"service_id"`
	Sum       uint `csv:"sum"`
}

func GetServicesInfoFromDB(input GetServicesInfoInput) (*[]ServiceInfo, error) {
	var serviceInfo []ServiceInfo
	err := config.DB.Raw("select service_id, sum(price) from orders where extract(year from orders.\"timestamp\") = " + strconv.Itoa(int(input.Year)) +
		" and extract(month from orders.\"timestamp\") = " + strconv.Itoa(int(input.Month)) + " and orders.is_completed = true group by service_id").
		Scan(&serviceInfo).Error
	if err != nil {
		return nil, err
	}
	return &serviceInfo, err
}

func CreateServicesInfoScv(serviceInfo []ServiceInfo) (*os.File, error) {
	servicesFile, err := os.OpenFile("./static/services.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return servicesFile, err
	}
	defer servicesFile.Close()

	err = gocsv.MarshalFile(&serviceInfo, servicesFile)
	if err != nil {
		return servicesFile, err
	}
	return servicesFile, err
}
