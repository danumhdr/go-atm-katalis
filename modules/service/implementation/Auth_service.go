package implementation

import (
	"fmt"
	"go-atm/modules/repository"
	"go-atm/modules/repository/implementation"
	"go-atm/modules/service"
)

type AuthService struct {
	Repository repository.AuthInterfaceRepository
}

func NewAuthService(authRepo repository.AuthInterfaceRepository) service.AuthInterfaceService {
	return &AuthService{Repository: authRepo}
}

func (service *AuthService) Login(username string) (bool, string) {
	var transRepo repository.TransactionInterfaceRepository
	result := implementation.NewAuthRepository().CheckUser(username)
	if !result {
		createUser := implementation.NewAuthRepository().CreateUser(username)
		if createUser != nil {
			return false, ""
		}
		_, err := implementation.NewTransactionRepository().CreateDeposit(username)
		if err != nil {
			return false, ""
		}
		return true, username
	}
	fmt.Printf("Hello, %s! \n", username)
	NewTransactionService(transRepo).CheckDeposit(username)
	return result, username
}

func (service *AuthService) Logout() {

}
