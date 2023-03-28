package repository

import (
	"github.com/jmoiron/sqlx"
	"todoApp/internal/model"
)

type ListRepository struct {
	db *sqlx.DB
}

func NewToDoListRepository(db *sqlx.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) AddList(userId int64, list model.ToDoList) (int64, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return -1, err
	}

	var listId int64

	addListQuery := "INSERT into todo_list (title, description) VALUES ($1, $2) RETURNING id"

	row := tx.QueryRowx(addListQuery, list.Title, list.Description)

	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return -1, err
	}

	addListIdToUsersList := "insert into users_lists (user_id, list_id) VALUES ($1, $2)"

	_, err = tx.Exec(addListIdToUsersList, userId, listId)
	if err != nil {
		tx.Rollback()
		return -1, nil
	}

	return listId, tx.Commit()
}
