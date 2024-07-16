package domain

import "errors"

var ErrNotFound = errors.New("usuário não encontrado")

type UserRepository[T any] interface {
	Create(model *T) error
	Delete(id uint) error
	FindByID(id uint) (*T, error)
	Update(model *T) error
}
