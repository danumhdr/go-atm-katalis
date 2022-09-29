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

func (transService *TransactionService) UpdateDeposit(username string, amount int) {
	result := implementation.NewTransactionRepository().CheckDeposit(username)
	newUserDeposit := model.DepositModel{User_deposit: username, Total_deposit: (result.Total_deposit + amount)}
	result = implementation.NewTransactionRepository().UpdateDeposit(newUserDeposit)
	transService.CheckDeposit(result.User_deposit)
}
func (transService *TransactionService) CheckDebt() {

}
func (transService *TransactionService) Transfer() {

}
func (transService *TransactionService) Withdraw() {

}
