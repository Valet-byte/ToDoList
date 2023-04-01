package repository

import (
	"github.com/jmoiron/sqlx"
	"todoApp/internal/model"
)

type PostgresItemRepository struct {
	db *sqlx.DB
}

func NewPostgresItemRepository(db *sqlx.DB) *PostgresItemRepository {
	return &PostgresItemRepository{
		db: db,
	}
}

func (r *PostgresItemRepository) AddItemList(listId int64, item model.ToDoItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int64

	createItemQuery := "INSERT INTO todo_item (title, description) values ($1, $2) RETURNING id"

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	createItemListQuery := "INSERT INTO items_lists (item_id, list_id) VALUES ($1, $2)"

	_, err = tx.Exec(createItemListQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return itemId, tx.Commit()
}

func (r *PostgresItemRepository) FindAll(userId, listId int64) ([]model.ToDoItem, error) {

	var items []model.ToDoItem

	selectAllItems := "SELECT * FROM todo_item ti " +
		"INNER JOIN items_lists il on ti.id = il.item_id " +
		"INNER JOIN users_lists ul on il.list_id = ul.list_id " +
		"WHERE il.list_id = $1 AND ul.user_id = $2"

	err := r.db.Select(&items, selectAllItems, listId, userId)
	return items, err
}
