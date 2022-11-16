package models

import (
	"time"
)

type UsersTransfer struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	UserSenderID    uint      `json:"user_sender_id"`
	UserGetterID    uint      `json:"user_getter_id"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
}
