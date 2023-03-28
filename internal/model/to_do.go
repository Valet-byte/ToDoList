package model

type ToDoList struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int64 `json:"id"`
	ListId int64 `json:"listId"`
	UserId int64 `json:"userId"`
}

type ToDoItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type ItemList struct {
	Id     int64 `json:"id"`
	ListId int64 `json:"listId"`
	ItemId int64 `json:"itemId"`
}
