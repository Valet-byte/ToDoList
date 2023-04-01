package repository

import (
	"github.com/jmoiron/sqlx"
	"todoApp/internal/model"
)

type AuthorizationRepository interface {
	AddUser(user model.User) (int64, error)
	GetUser(username, password string) (model.User, error)
}

type TodoListRepository interface {
	AddList(userId int64, list model.ToDoList) (int64, error)
	FindAll(userId int64) ([]model.ToDoList, error)
	FindById(userId, listId int64) (model.ToDoList, error)
	Delete(userId, listId int64) error
	UpdateList(userId, listId int64, input model.UpdateListInput) error
}

type ItemSRepository interface {
	AddItemList(listId int64, item model.ToDoItem) (int64, error)
	FindAll(userId, listId int64) ([]model.ToDoItem, error)
}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	ItemSRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthRepository(db),
		TodoListRepository:      NewToDoListRepository(db),
		ItemSRepository:         NewPostgresItemRepository(db),
	}
}
