package service

import (
	"github.com/sirupsen/logrus"
	"todoApp/internal/model"
	"todoApp/internal/repository"
)

type ItemService struct {
	repo     repository.ItemSRepository
	listRepo repository.TodoListRepository
}

func NewItemService(repo *repository.ItemSRepository, listRepo *repository.TodoListRepository) *ItemService {
	return &ItemService{
		repo:     *repo,
		listRepo: *listRepo,
	}
}

func (s *ItemService) CreateItemList(userId, listId int64, item model.ToDoItem) (int64, error) {
	_, err := s.listRepo.FindById(userId, listId)
	if err != nil {
		logrus.Debugf("List with listId : %s and userId : %s not found!", listId, userId)
		return -1, err
	}
	return s.repo.AddItemList(listId, item)
}
