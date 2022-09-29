package controller

import (
	"go-atm/config"
	"go-atm/model"
	"go-atm/modules/repository"
	"go-atm/modules/service/implementation"
)

var authRepo repository.AuthInterfaceRepository

func Login(username string) {
	validateUser, userID := implementation.NewAuthService(authRepo).Login(username)
	if validateUser {
		config.Session = &model.UserModel{Name: userID}
	}
}

func Logout() {
	config.Session = &model.UserModel{}
}
