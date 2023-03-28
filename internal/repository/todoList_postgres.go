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
func (r *ListRepository) FindAll(userId int64) ([]model.ToDoList, error) {
	var lists []model.ToDoList
	query := "SELECT tl.id, tl.title, tl.description from todo_list tl INNER JOIN users_lists ul on tl.id = ul.list_id WHERE ul.user_id = $1"

	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ListRepository) FindById(userId, listId int64) (model.ToDoList, error) {
	var list model.ToDoList
	query := "SELECT tl.id, tl.title, tl.description from todo_list tl INNER JOIN users_lists ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2"

	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
