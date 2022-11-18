package utils

import (
	"fmt"
	"github.com/Owner-maker/microservice-for-working-with-user-balance/config"
	"strconv"
	"time"
)

const (
	selfIncomeDescription                    string = "replenishment of the balance"
	transferToUserDescription                string = "replenishment of balance to user"
	selfIncomeFromUserDescription            string = "replenishment of the balance from the user"
	paymentOfServiceDescription              string = "payment of service"
	selfIncomeWhenServiceCanceledDescription string = "replenishment of the balance when canceling the service"
)

type Pagination struct {
	Limit uint   `json:"limit"`
	Page  uint   `json:"page"`
	Sort  string `json:"sort"`
}

type UserTransaction struct {
	UserID          uint      `json:"user_id"`
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
	AnotherUserID   uint      `json:"another_user_id"`
	ServiceID       uint      `json:"service_id"`
	IsCompleted     bool      `json:"is_completed"`
}

type UserFormattedTransaction struct {
	IncomingBalance uint      `json:"incoming_balance"`
	OutgoingBalance uint      `json:"outgoing_balance"`
	Timestamp       time.Time `json:"timestamp"`
	MoneyValue      uint      `json:"money_value"`
	Description     string    `json:"description"`
}

func GetPaginatedUserTransactions(userID uint, pagination Pagination) (*[]UserFormattedTransaction, error) {
	var userTransactions []UserTransaction
	var userFormattedTransactions []UserFormattedTransaction
	offset := (pagination.Page - 1) * pagination.Limit
	err := config.DB.Raw(
		"select incoming_balance, outgoing_balance,\"timestamp\", money_value, " +
			"Null as service_id,Null as is_completed, 0 as another_user_id from self_incomes where user_id = " + strconv.Itoa(int(userID)) +
			" union select incoming_balance, outgoing_balance,\"timestamp\",price as money_value, " +
			"service_id,is_completed, 0 as another_user_id from orders  where user_id = " + strconv.Itoa(int(userID)) +
			" union select incoming_balance, outgoing_balance,\"timestamp\", money_value, " +
			"Null as service_id,Null as is_completed, another_user_id from users_transfers where user_id = " + strconv.Itoa(int(userID)),
	).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Scan(&userTransactions).Error
	if err != nil {
		return nil, err
	}
	var description string
	for _, v := range userTransactions {
		currTransaction := UserFormattedTransaction{
			IncomingBalance: v.IncomingBalance,
			OutgoingBalance: v.OutgoingBalance,
			Timestamp:       v.Timestamp,
			MoneyValue:      v.MoneyValue,
		}
		if v.AnotherUserID != 0 {
			anotherUser := GetUser(v.AnotherUserID)
			if v.IncomingBalance < v.OutgoingBalance {
				description = fmt.Sprintf("%s %s", selfIncomeFromUserDescription, anotherUser.Nickname)
			} else if v.IncomingBalance > v.OutgoingBalance {
				description = fmt.Sprintf("%s %s", transferToUserDescription, anotherUser.Nickname)
			}
		} else if v.ServiceID != 0 {
			if v.IsCompleted == true {
				description = fmt.Sprintf("%s #%s", paymentOfServiceDescription, strconv.Itoa(int(v.ServiceID)))
			} else if v.OutgoingBalance > v.IncomingBalance {
				description = fmt.Sprintf("%s #%s", selfIncomeWhenServiceCanceledDescription, strconv.Itoa(int(v.ServiceID)))
			}
		} else {
			description = selfIncomeDescription
		}
		currTransaction.Description = description
		userFormattedTransactions = append(userFormattedTransactions, currTransaction)
		fmt.Println(userFormattedTransactions)
	}
	return &userFormattedTransactions, nil
}
