package repository

type Repository[T any] interface {
	Create(T) error
	Read() ([]T, error)
}