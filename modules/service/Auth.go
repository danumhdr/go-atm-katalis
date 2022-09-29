package service

type AuthInterfaceService interface {
	Login(username string) (bool, string)
	Logout()
}
