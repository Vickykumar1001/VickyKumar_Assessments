package service

import (
	"A2_EMS_PROJECT/config"
	"A2_EMS_PROJECT/model"
	"A2_EMS_PROJECT/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (service *UserService) Signup(user *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	return service.UserRepo.CreateUser(user)
}

func (service *UserService) Login(username, password string) (string, error) {
	user, err := service.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := config.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
