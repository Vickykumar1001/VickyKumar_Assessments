package service

import "A1_BMS_PROJECT/repository"

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{AuthRepo: repo}
}

func (s *AuthService) Signup(username, password string) error {
	return s.AuthRepo.CreateUser(username, password)
}
