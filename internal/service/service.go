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
	DeleteList(userId, listId int64) error
	Update(userId, listId int64, input model.UpdateListInput) error
}

type ToDoItemService interface {
	CreateItemList(userId, listId int64, item model.ToDoItem) (int64, error)
}

type Service struct {
	AuthorizationService
	TodoListService
	ToDoItemService
}

func NewService(repository *repository.Repository, jwtKey, passwordSalt string, tokenExpiresAt int) *Service {
	return &Service{
		AuthorizationService: NewAuthService(&repository.AuthorizationRepository, jwtKey, passwordSalt, tokenExpiresAt),
		TodoListService:      NewListService(&repository.TodoListRepository),
		ToDoItemService:      NewItemService(&repository.ItemSRepository, &repository.TodoListRepository),
	}
}
