package repository

type AuthInterfaceRepository interface {
	CheckUser(username string) bool
	CreateUser(username string) error
}
