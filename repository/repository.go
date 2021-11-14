package repository

type Repository interface {
	Insert(data interface{}) error
}
