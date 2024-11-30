package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (c *Repository[T]) Create(tx *gorm.DB, entity T) error {
	return tx.Create(entity).Error
}
