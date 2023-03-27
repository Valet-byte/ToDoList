package service

import (
	"crypto/sha512"
	"fmt"
	"todoApp/internal/model"
	"todoApp/internal/repository"
)

const salt = "asfndi3(*(39@!122kdsek+_09"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int64, error) {
	user.Password = s.createHashPassword(user.Password)
	return s.repo.AuthorizationRepository.AddUser(user)
}

func (s *AuthService) createHashPassword(pass string) string {
	sha := sha512.New()
	sha.Write([]byte(pass))

	return fmt.Sprintf("%x", sha.Sum([]byte(salt)))
}
