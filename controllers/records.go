package controllers

import "time"

const (
	selfIncomeDescription                    string = "replenishment of the balance"
	transferToUserDescription                string = "replenishment of balance to user"
	selfIncomeFromUserDescription            string = "replenishment of the balance from the user"
	paymentOfServiceDescription              string = "payment of service"
	selfIncomeWhenServiceCanceledDescription string = "replenishment of the balance when canceling the service"
)

type UserTransaction struct {
	UserID          uint      `json:"user_id"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
	Description     string    `json:"description"`
}
