package implementation

import (
	"go-atm/config"
	"go-atm/model"
	"go-atm/modules/repository"
	"log"
)

type AuthRepository struct {
}

func NewAuthRepository() repository.AuthInterfaceRepository {
	return &AuthRepository{}
}

func (authRepo *AuthRepository) CheckUser(username string) bool {
	var user model.UserModel
	getRecord := config.DB.Table("tbl_users").Where("name = ?", username).First(&user)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return false
	}
	return true
}

func (authRepo *AuthRepository) CreateUser(username string) error {
	user := model.UserModel{Name: username}
	getRecord := config.DB.Table("tbl_users").Create(&user)
	if getRecord.Error != nil {
		log.Println(getRecord.Error)
		return getRecord.Error
	}
	return nil
}
