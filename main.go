package main

import (
	"fmt"
	"myapp/constants"
	"myapp/models"
	"time"
)

func main() {
	fmt.Println("hello world")
	//we first create digital wallet
	dw, err := models.NewDigitalWallet()
	if err != nil {
		panic("digital wallet could not be created.")
	}
	// we will create users
	user1, err := models.NewUser("Pranav", "Password123", 26, "abc@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user2, err := models.NewUser("Prateek", "Password123", 28, "abc1@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// add users in our db
	go dw.CreateUser(*user1)
	go dw.CreateUser(*user2)

	//now we create accounts
	acc1, err := models.NewAccount(500, constants.USD, user1.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	acc2, err := models.NewAccount(1000, constants.INR, user2.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//add accounts in our db
	go dw.CreateAccount(*acc1)
	go dw.CreateAccount(*acc2)

	time.Sleep(2 * time.Second) //adding this to be sure that we have account created before doing transaction

	// we make transaction
	go dw.TransferFunds(acc1.ID, acc2.ID, 100) // we transferred 100$ so as per ration acc2 should receive 2000 INR
	go dw.TransferFunds(acc2.ID, acc1.ID, 500) // we transferred 500INR so as per ration acc1 should receive 25$

	time.Sleep(10 * time.Second)

	//we check if transaction was successfull
	regAcc1, _ := dw.GetAccount(acc1.ID)
	regAcc2, _ := dw.GetAccount(acc2.ID)

	fmt.Println("Balance in Account 1 ==>", regAcc1.Balance)
	fmt.Println("Balance in Account 2 ==>", regAcc2.Balance)

}
