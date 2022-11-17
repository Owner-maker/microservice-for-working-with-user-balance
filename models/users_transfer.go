package models

import (
	"time"
)

type UsersTransfer struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	UserID          uint      `json:"user_id"`
	AnotherUserID   uint      `json:"another_user_id"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
}
