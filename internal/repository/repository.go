package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (c *Repository[T]) Create(tx *gorm.DB, entity *T) error {
	return tx.Create(entity).Error
}

func (c *Repository[T]) BulkCreate(tx *gorm.DB, entities []*T) error {
	return tx.Create(entities).Error
}

func (c *Repository[T]) Update(tx *gorm.DB, entity *T) error {
	return tx.Save(entity).Error
}

func (c *Repository[T]) Delete(tx *gorm.DB, entity *T) error {
	return tx.Delete(entity).Error
}

func (c *Repository[T]) FindById(tx *gorm.DB, entity *T, id any) error {
	return tx.Where("id = ?", id).Take(entity).Error
}
