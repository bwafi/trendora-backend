package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID       string `gorm:"column:id;type:uuid;primaryKey"`
	ParentId string `gorm:"column:parent_id"`
	Name     string `gorm:"column:name;type:varchar(100);not null"`
	Slug     string `gorm:"column:slug;type:varchar(100);not null"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}