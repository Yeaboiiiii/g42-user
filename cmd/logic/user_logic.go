package logic

import (
	"g42-user/repositories"
	"g42-user/utils"
)

type UserLogic struct {
	userRepo *repositories.UserRepository
}

func NewUserLogic(userRepo *repositories.UserRepository) *UserLogic {
	return &UserLogic{userRepo: userRepo}
}

func (l *UserLogic) Login(email, password string) (string, repositories.User, error) {
	user, err := l.userRepo.FindByEmail(email)
	if err != nil {
		return "", repositories.User{}, err
	}

	if !l.userRepo.ValidatePassword(email, password) {
		return "", repositories.User{}, nil
	}

	token, err := utils.GenerateToken(email)
	if err != nil {
		return "", repositories.User{}, err
	}

	return token, *user, nil
}

func (l *UserLogic) Register(user *repositories.User) error {
	return l.userRepo.CreateUser(user)
}

func (l *UserLogic) GetUserByEmail(email string) (*repositories.User, error) {
	return l.userRepo.FindByEmail(email)
}

func (l *UserLogic) GetUserByID(id string) (*repositories.User, error) {
	return l.userRepo.FindByID(id)
}
