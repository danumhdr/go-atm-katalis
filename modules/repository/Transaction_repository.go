package repository

import "go-atm/model"

type TransactionInterfaceRepository interface {
	CreateDeposit(username string) (model.DepositModel, error)
	CheckDeposit(username string) model.DepositModel
	UpdateDeposit(userDeposit model.DepositModel) model.DepositModel
	Transfer(transferFrom model.DepositModel, transferTo model.DepositModel) model.DepositModel
	CreateDebt(userDebt model.DebtModel) ([]model.DebtModel, error)
	UpdateDebt(userDebt model.DebtModel) []model.DebtModel
	CheckDebt(userDebt model.DebtModel) []model.DebtModel
}
