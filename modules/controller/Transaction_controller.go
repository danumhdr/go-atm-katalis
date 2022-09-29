package controller

import (
	"fmt"
	"go-atm/config"
	"go-atm/modules/repository"
	"go-atm/modules/service/implementation"
	"strconv"
)

var transRepo repository.TransactionInterfaceRepository

func Deposit(amount string) {
	username := config.Session.Name
	if username != "" {
		amounts, _ := strconv.Atoi(amount)
		implementation.NewTransactionService(transRepo).UpdateDeposit(username, amounts)
	}
}

func Transfer() {
	fmt.Println("Transfer")
}

func CheckAmount() {
	fmt.Println("CheckAmount")
}

func Withdraw(amount string) {
	username := config.Session.Name
	if username != "" {
		amounts, _ := strconv.Atoi(amount)
		implementation.NewTransactionService(transRepo).UpdateDeposit(username, -amounts)
	}
}
