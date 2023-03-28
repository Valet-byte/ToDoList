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
	CreateList(userId int64, list model.ToDoList) (int64, error)
	GetAll(userId int64) ([]model.ToDoList, error)
	GetById(userId, listId int64) (model.ToDoList, error)
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
		AuthorizationService: NewAuthService(&repository.AuthorizationRepository, jwtKey, passwordSalt, tokenExpiresAt),
		TodoListService:      NewListService(&repository.TodoListRepository),
	}
}
