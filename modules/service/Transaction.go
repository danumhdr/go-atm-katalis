package service

type TransactionInterfaceService interface {
	CheckDeposit(username string)
	UpdateDeposit(username string, amount int)
	Transfer(transferFrom string, transferTo string, amount int)
}
