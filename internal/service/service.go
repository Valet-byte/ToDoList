package service

import (
	"todoApp/internal/model"
	"todoApp/internal/repository"
)

type AuthorizationService interface {
	CreateUser(user model.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type TodoListService interface {
}

type ItemService interface {
}

type Service struct {
	AuthorizationService
	TodoListService
	ItemService
}

func NewService(repository *repository.Repository, jwtKey, passwordSalt string, tokenExpiresAt int) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repository, jwtKey, passwordSalt, tokenExpiresAt),
	}
}
