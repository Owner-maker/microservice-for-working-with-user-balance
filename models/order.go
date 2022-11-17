package models

import (
	"time"
)

type Order struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	UserID          uint      `json:"user_id"`
	ServiceID       uint      `json:"service_id"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	IsCompleted     bool      `json:"is_completed" gorm:"default:false"`
	Price           uint      `json:"price"`
}
