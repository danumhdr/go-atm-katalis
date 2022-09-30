package repository

import "go-atm/model"

type TransactionInterfaceRepository interface {
	CreateDeposit(username string) (model.DepositModel, error)
	CheckDeposit(username string) model.DepositModel
	UpdateDeposit(userDeposit model.DepositModel) model.DepositModel
	Transfer(transferFrom string, transferTo string, amount int) (model.DepositModel, model.DepositModel)
	CreateDebt(userDebt model.DebtModel) (model.DebtModel, error)
	UpdateDebt(userDebt model.DebtModel) error
	DeleteDebt(transferFrom model.DebtModel) error
	CheckAllDebt(username string) []model.DebtModel
	CheckDetailDebt(username string, userDebtFor string) model.DebtModel
}
