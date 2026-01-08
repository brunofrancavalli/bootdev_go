package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(cust *customer, trans transaction) error {
	switch trans.transactionType {
	case "deposit":
		cust.balance += trans.amount
		return nil
	case "withdrawal":
		if cust.balance > trans.amount {
			cust.balance -= trans.amount
			return nil
		} else {
			return errors.New("insufficient funds")
		}
	default:
		return errors.New("unknown transaction type")
	}
}
