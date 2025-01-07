package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID          string  `gorm:"column:id;type:uuid;primaryKey"`
	ProductId   string  `gorm:"column:product_id;type:varchar(255)"`
	SKU         string  `gorm:"column:sku;type:varchar(100);not null"`
	ColorName   string  `gorm:"column:color_name;type:varchar(50)"`
	Weight      float32 `gorm:"column:weight;type:decimal(8,2)"`
	IsAvailable bool    `gorm:"column:is_available;type:bool"`

	Product Product `gorm:"foreignKey:product_id;references:id"`
}

func (c *ProductVariant) TableName() string {
	return "product_variants"
}

func (c *ProductVariant) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
