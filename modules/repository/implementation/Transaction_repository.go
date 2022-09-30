package implementation

import (
	"go-atm/config"
	"go-atm/model"
	"go-atm/modules/repository"
	"log"
)

type TransactionRepository struct {
}

func NewTransactionRepository() repository.TransactionInterfaceRepository {
	return &TransactionRepository{}
}

func (transRepo *TransactionRepository) CreateDeposit(username string) (model.DepositModel, error) {
	userDeposit := model.DepositModel{User_deposit: username, Total_deposit: 0}
	getRecord := config.DB.Table("tbl_deposits").Create(&userDeposit)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDeposit, getRecord.Error
	}
	return userDeposit, nil
}

func (transRepo *TransactionRepository) CheckDeposit(username string) model.DepositModel {
	var userDeposit model.DepositModel
	getRecord := config.DB.Table("tbl_deposits").Where("user_deposit = ?", username).First(&userDeposit)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDeposit
	}
	return userDeposit
}

func (transRepo *TransactionRepository) UpdateDeposit(userDeposit model.DepositModel) model.DepositModel {
	var userDeposits model.DepositModel

	getRecord := config.DB.Table("tbl_deposits").Where("user_deposit = ?", userDeposit.User_deposit).Update("total_deposit", userDeposit.Total_deposit)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDeposit
	}
	userDeposits = transRepo.CheckDeposit(userDeposit.User_deposit)
	return userDeposits
}

func (transRepo *TransactionRepository) Transfer(transferFrom string, transferTo string, amount int) (model.DepositModel, model.DepositModel) {

	userFromDeposits := transRepo.CheckDeposit(transferFrom)
	userFromDeposits.Total_deposit -= amount
	getTransferFrom := transRepo.UpdateDeposit(userFromDeposits)

	userToDeposits := transRepo.CheckDeposit(transferTo)
	userToDeposits.Total_deposit += amount
	getTransferTo := transRepo.UpdateDeposit(userToDeposits)

	return getTransferFrom, getTransferTo
}

func (transRepo *TransactionRepository) CreateDebt(transferFrom model.DebtModel) (model.DebtModel, error) {
	var userDebt model.DebtModel
	getRecord := config.DB.Table("tbl_users_debt").Create(&transferFrom)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDebt, getRecord.Error
	}
	return userDebt, nil
}

func (transRepo *TransactionRepository) UpdateDebt(transferFrom model.DebtModel) error {
	if transferFrom.Amount_debts == 0 {
		return transRepo.DeleteDebt(transferFrom)
	}
	getRecord := config.DB.Table("tbl_users_debt").Where("user_debts = ? and user_debts_for = ?", transferFrom.User_debts, transferFrom.User_debts_for).Update("amount_debts", transferFrom.Amount_debts)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return getRecord.Error
	}
	return nil
}

func (transRepo *TransactionRepository) DeleteDebt(transferFrom model.DebtModel) error {
	getRecord := config.DB.Table("tbl_users_debt").Where("user_debts = ? and user_debts_for = ?", transferFrom.User_debts, transferFrom.User_debts_for).Delete(&transferFrom)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return getRecord.Error
	}
	return nil
}

func (transRepo *TransactionRepository) CheckAllDebt(username string) []model.DebtModel {
	var userDebts []model.DebtModel
	getRecord := config.DB.Table("tbl_users_debt").Where("user_debts = ? and amount_debts > 0", username).Order("amount_debts DESC").Scan(&userDebts)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDebts
	}
	return userDebts
}

func (transRepo *TransactionRepository) CheckDetailDebt(username string, userDebtFor string) model.DebtModel {
	var userDebts model.DebtModel
	getRecord := config.DB.Table("tbl_users_debt").Where("user_debts = ? and user_debts_for = ?", username, userDebtFor).Order("amount_debts DESC").Scan(&userDebts)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDebts
	}
	return userDebts
}
