package models

import (
	con "myapp/constants"
	customError "myapp/error"
	"myapp/util"
	"time"
)

type Account struct {
	Base
	Balance       float64
	AccountNumber int
	Currency      con.Currency
	UserID        int
	Transactions  []*Transaction
}

func NewAccount(bal float64, cur con.Currency, userId int) (*Account, error) {
	var acc = Account{
		Base: Base{
			ID:        util.GetUniqueAccountID(),
			CreatedAt: time.Now(),
		},
		Balance:       bal,
		AccountNumber: util.GetUniqueAccountID(),
		Currency:      cur,
		UserID:        userId,
	}

	err := acc.Validate()
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (a *Account) Validate() error {

	if a.Balance < 0 {
		return customError.NewValidationError("Balance can't be -ve.")
	}

	if !a.Currency.IsValidCurrency() {
		return customError.NewValidationError("Currency is Invalid.")
	}

	return nil
}
