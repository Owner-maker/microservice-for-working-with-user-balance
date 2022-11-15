package models

import "time"

type SelfIncome struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
}
