package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todoApp/internal/model"
	"todoApp/internal/repository/db"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) AddUser(user model.User) (int64, error) {
	var id int64
	query := fmt.Sprintf("insert into %s (name, username, password) values ($1, $2, $3) RETURNING id", db.UserTable)
	row := r.db.QueryRowx(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
