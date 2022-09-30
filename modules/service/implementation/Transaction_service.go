package implementation

import (
	"fmt"
	"go-atm/model"
	"go-atm/modules/repository"
	"go-atm/modules/repository/implementation"
	"go-atm/modules/service"
)

type TransactionService struct {
	Repository repository.TransactionInterfaceRepository
}

func NewTransactionService(transRepo repository.TransactionInterfaceRepository) service.TransactionInterfaceService {
	return &TransactionService{Repository: transRepo}
}

func (transService *TransactionService) CheckDeposit(username string) {
	result := implementation.NewTransactionRepository().CheckDeposit(username)
	fmt.Println("Your balance is $", result.Total_deposit)
}

func (transService *TransactionService) checkDebt(username string) {
	result := implementation.NewTransactionRepository().CheckAllDebt(username)
	for _, debt := range result {
		fmt.Printf("Owed %d to %s\n", debt.Amount_debts, debt.User_debts_for)
	}
}

func (transService *TransactionService) UpdateDeposit(username string, amount int) {
	result := implementation.NewTransactionRepository().CheckDeposit(username)
	newUserDeposit := model.DepositModel{User_deposit: username, Total_deposit: (result.Total_deposit + amount)}
	result = implementation.NewTransactionRepository().UpdateDeposit(newUserDeposit)
	getAllDebt := implementation.NewTransactionRepository().CheckAllDebt(username)
	for _, debt := range getAllDebt {
		amountDebt := debt.Amount_debts
		if result.Total_deposit < amountDebt {
			amountDebt -= result.Total_deposit
			result.Total_deposit = 0
		} else {
			result.Total_deposit -= amountDebt
		}
		transService.Transfer(username, debt.User_debts_for, amountDebt)
		checkDebt := model.DebtModel{User_debts: username, User_debts_for: debt.User_debts_for, Amount_debts: debt.Amount_debts - amountDebt}
		_ = implementation.NewTransactionRepository().UpdateDebt(checkDebt)
		if result.Total_deposit == 0 {
			break
		}
	}
	transService.CheckDeposit(result.User_deposit)
}

func (transService *TransactionService) Transfer(transferFrom string, transferTo string, amount int) {
	result := implementation.NewTransactionRepository().CheckDeposit(transferFrom)
	var debt int = 0
	if result.Total_deposit < amount {
		checkDebt := implementation.NewTransactionRepository().CheckDetailDebt(transferFrom, transferTo)
		if checkDebt.User_debts_for != "" {
			debt = amount - result.Total_deposit
			checkDebt.Amount_debts += debt
			if debt == 0 {
				checkDebt.Amount_debts = 0
			}
			_ = implementation.NewTransactionRepository().UpdateDebt(checkDebt)
		} else {
			debt = amount - result.Total_deposit
			checkDebt = model.DebtModel{User_debts: transferFrom, User_debts_for: transferTo, Amount_debts: debt}
			_, _ = implementation.NewTransactionRepository().CreateDebt(checkDebt)
		}
		depositFrom, depositTo := implementation.NewTransactionRepository().Transfer(transferFrom, transferTo, result.Total_deposit)
		fmt.Printf("Transferred %d to %s\n", amount, depositTo.User_deposit)
		transService.CheckDeposit(depositFrom.User_deposit)
		fmt.Printf("Owed %d to %s\n", checkDebt.Amount_debts, depositTo.User_deposit)
		result.Total_deposit = 0
	} else {
		result.Total_deposit -= amount
		depositFrom, depositTo := implementation.NewTransactionRepository().Transfer(transferFrom, transferTo, amount)
		fmt.Printf("Transferred %d to %s\n", amount, depositTo.User_deposit)
		transService.CheckDeposit(depositFrom.User_deposit)
	}
}
