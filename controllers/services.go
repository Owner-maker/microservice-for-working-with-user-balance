package controllers

type ReserveMoneyForTheService struct {
	UserID    uint `json:"user_id" binding:"required"`
	ServiceID uint `json:"service_id" binding:"required"`
}
