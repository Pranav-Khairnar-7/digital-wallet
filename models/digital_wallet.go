package models

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type DigitalWallet struct {
	Base
	Users        map[int]*User
	Accounts     map[int]*Account
	Transactions map[int][]*Transaction
	mu           sync.Mutex
}

func NewDigitalWallet() (*DigitalWallet, error) {

	var wallet = DigitalWallet{
		Base: Base{
			ID:        999999,
			CreatedAt: time.Now(),
		},
		Users:        make(map[int]*User),
		Accounts:     make(map[int]*Account),
		Transactions: map[int][]*Transaction{},
	}

	return &wallet, nil
}

func (d *DigitalWallet) GetRegisteredUser(userId int) (*User, error) {
	user, ok := d.Users[userId]

	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (d *DigitalWallet) GetAccount(accId int) (*Account, error) {
	acc, ok := d.Accounts[accId]

	if !ok {
		return nil, errors.New("account not found")
	}

	return acc, nil
}

func (d *DigitalWallet) CreateUser(user User) error {

	d.mu.Lock()
	defer d.mu.Unlock()

	d.Users[user.ID] = &user

	return nil
}

func (d *DigitalWallet) CreateAccount(acc Account) error {

	d.mu.Lock()
	defer d.mu.Unlock()

	d.Accounts[acc.ID] = &acc

	return nil
}

func (d *DigitalWallet) TransferFunds(fromAccId, toAccId int, amount float64) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Println("transfer funds called ==>", fromAccId, toAccId, amount)
	//get from account
	fromAcc, ok := d.Accounts[fromAccId]
	if !ok {
		fmt.Println("from account not found")
		return errors.New("from account not found")
	}

	//get to account
	toAcc, ok := d.Accounts[toAccId]
	if !ok {
		fmt.Println("to account not found")
		return errors.New("to account not found")
	}

	if fromAcc.Balance < amount {
		fmt.Println("insufficient balance")
		return errors.New("insufficient balance")
	}

	// amount to be credited in to acc
	var amountToCredit float64

	if toAcc.Currency != fromAcc.Currency {
		amountToCredit = toAcc.Currency.CurrencyConverter(fromAcc.Currency, toAcc.Currency, amount)
	} else {
		amountToCredit = amount
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amountToCredit

	//create new transaction
	tran, err := NewTransaction(fromAccId, toAccId, amount, fromAcc.Currency)

	if err != nil {
		return err
	}

	// we will set this transaction for both accounts
	d.Transactions[fromAccId] = append(d.Transactions[fromAccId], tran)
	d.Transactions[toAccId] = append(d.Transactions[toAccId], tran)

	// we will add it in account map transaction array as well
	fromAcc.Transactions = append(fromAcc.Transactions, tran)
	toAcc.Transactions = append(toAcc.Transactions, tran)

	fmt.Println("This is from acc in tran =>", fromAcc.Balance)
	fmt.Println("This is to acc in tran =>", toAcc.Balance)

	return nil
}
