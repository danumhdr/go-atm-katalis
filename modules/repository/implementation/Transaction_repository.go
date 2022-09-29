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
	//check punya hutang/debts atau nggak

	//kalau punya dia harus transfer ke yg diutangin ke yang paling banyak

	//lihat sisa deposit, kalau masi ada baru update

	getRecord := config.DB.Table("tbl_deposits").Where("user_deposit = ?", userDeposit.User_deposit).Update("total_deposit", userDeposit.Total_deposit)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return userDeposit
	}
	userDeposits = transRepo.CheckDeposit(userDeposit.User_deposit)
	return userDeposits
}

func (transRepo *TransactionRepository) Transfer(transferFrom model.DepositModel, transferTo model.DepositModel) model.DepositModel {
	//check deposit saldo

	//kalau deposit < transfer

	//1 transfer all deposit ke transfer tujuan

	//2 buat record debt untuk user yg melakukan transfer
	return model.DepositModel{}
}

func (transRepo *TransactionRepository) CreateDebt(transferFrom model.DebtModel) ([]model.DebtModel, error) {
	return []model.DebtModel{}, nil
}

func (transRepo *TransactionRepository) UpdateDebt(transferFrom model.DebtModel) []model.DebtModel {
	return []model.DebtModel{}
}

func (transRepo *TransactionRepository) CheckDebt(userDebt model.DebtModel) []model.DebtModel {
	return []model.DebtModel{}
}
