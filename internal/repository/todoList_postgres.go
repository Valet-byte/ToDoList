package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *ListRepository) Delete(userId, listId int64) error {
	query := "delete from todo_list tl USING users_lists ul where tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2"

	row, err := r.db.Exec(query, userId, listId)
	in, er := row.RowsAffected()
	if er != nil {
		return er
	}

	if in == 0 {
		return errors.New("not delete list")
	}
	return err
}

func (r *ListRepository) UpdateList(userId, listId int64, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	var argId = 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE todo_list tl SET %s FROM users_lists ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		setQuery, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("Query : %s", query)
	logrus.Debugf("Args : %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
