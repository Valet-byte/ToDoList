package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
