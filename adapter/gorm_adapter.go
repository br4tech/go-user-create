package adapter

import (
	"errors"
	"fmt"

	"github.com/br4tech/go-user-create/domain"
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{
		db: db,
	}
}

func (r *GormRepository[T]) Create(model *T) error {
	if err := r.db.Create(model).Error; err != nil {
		return fmt.Errorf("erro ao criar registro: %w", err)
	}
	return nil
}

func (r *GormRepository[T]) FindByID(id uint) (*T, error) {
	var model T
	if err := r.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("erro ao buscar registro: %w", err)
	}
	return &model, nil
}

func (r *GormRepository[T]) Update(model *T) error {
	if err := r.db.Save(model).Error; err != nil {
		return fmt.Errorf("erro ao atualizar registro: %w", err)
	}
	return nil
}

func (r *GormRepository[T]) Delete(id uint) error {
	var model T
	if err := r.db.Delete(&model, id).Error; err != nil {
		return fmt.Errorf("erro ao excluir registro: %w", err)
	}
	return nil
}
