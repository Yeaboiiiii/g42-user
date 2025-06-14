package contracts

import "g42-user/cmd/repositories/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	ValidatePassword(email, password string) bool
}

type UserLogic interface {
	Login(email, password string) (string, models.User, error)
	Register(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}
