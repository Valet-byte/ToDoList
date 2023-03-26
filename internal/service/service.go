package service

import "todoApp/internal/repository"

type AuthorizationService interface {
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

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
