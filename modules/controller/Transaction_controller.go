package controller

import (
	"go-atm/config"
	"go-atm/modules/repository"
	"go-atm/modules/service/implementation"
	"strconv"
	"strings"
)

var transRepo repository.TransactionInterfaceRepository

func Deposit(amount string) {
	username := config.Session.Name
	if username != "" {
		amounts, _ := strconv.Atoi(amount)
		implementation.NewTransactionService(transRepo).UpdateDeposit(username, amounts)
	}
}

func Transfer(transferTo string, amount string) {
	username := config.Session.Name
	if strings.ToLower(username) != strings.ToLower(transferTo) {
		if username != "" {
			amounts, _ := strconv.Atoi(amount)
			implementation.NewTransactionService(transRepo).Transfer(username, transferTo, amounts)
		}
	}
}

func Withdraw(amount string) {
	username := config.Session.Name
	if username != "" {
		amounts, _ := strconv.Atoi(amount)
		implementation.NewTransactionService(transRepo).UpdateDeposit(username, -amounts)
	}
}
