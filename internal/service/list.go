package service

import (
	"todoApp/internal/model"
	"todoApp/internal/repository"
)

type ListService struct {
	repo repository.TodoListRepository
}

func NewListService(repo *repository.TodoListRepository) *ListService {
	return &ListService{repo: *repo}
}

func (s *ListService) CreateList(userId int64, list model.ToDoList) (int64, error) {
	return s.repo.AddList(userId, list)
}

func (s *ListService) GetAll(userId int64) ([]model.ToDoList, error) {
	return s.repo.FindAll(userId)
}

func (s *ListService) GetById(userId, listId int64) (model.ToDoList, error) {
	return s.repo.FindById(userId, listId)
}
