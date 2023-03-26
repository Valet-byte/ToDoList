package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
}

type TodoListRepository interface {
}

type ItemSRepository interface {
}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	ItemSRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
