package logic

import (
	"g42-user/cmd/logic/contracts"
	"g42-user/cmd/repositories/models"
	"g42-user/cmd/utils"
)

type UserLogic struct {
	userRepo contracts.UserRepository
}

func NewUserLogic(userRepo contracts.UserRepository) *UserLogic {
	return &UserLogic{userRepo: userRepo}
}

func (l *UserLogic) Login(email, password string) (string, models.User, error) {
	user, err := l.userRepo.FindByEmail(email)
	if err != nil {
		return "", models.User{}, err
	}

	if !l.userRepo.ValidatePassword(email, password) {
		return "", models.User{}, nil
	}

	token, err := utils.GenerateToken(email)
	if err != nil {
		return "", models.User{}, err
	}

	return token, *user, nil
}

func (l *UserLogic) Register(user *models.User) error {
	return l.userRepo.CreateUser(user)
}

func (l *UserLogic) GetUserByEmail(email string) (*models.User, error) {
	return l.userRepo.FindByEmail(email)
}

func (l *UserLogic) GetUserByID(id string) (*models.User, error) {
	return l.userRepo.FindByID(id)
}
