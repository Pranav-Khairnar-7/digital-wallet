package models

import (
	con "myapp/constants"
	"myapp/util"
	"time"
)

type Transaction struct {
	Base
	FromAccount     int
	ToAccount       int
	TransactionTime time.Time
	Amount          float64
	Currency        con.Currency
}

func NewTransaction(from int, to int, amount float64, cur con.Currency) (*Transaction, error) {

	return &Transaction{
		Base: Base{
			ID:        util.GetUniqueTransactionId(),
			CreatedAt: time.Now(),
		},
		FromAccount:     from,
		ToAccount:       to,
		TransactionTime: time.Now(),
		Amount:          amount,
		Currency:        cur,
	}, nil

}
