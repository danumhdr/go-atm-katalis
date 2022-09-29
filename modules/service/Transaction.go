package service

type TransactionInterfaceService interface {
	CheckDeposit(username string)
	UpdateDeposit(username string, amount int)
	CheckDebt()
	Transfer()
	Withdraw()
}
