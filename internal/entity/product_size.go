package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductSize struct {
	ID            string  `gorm:"column:id;type:uuid;primaryKey"`
	VariantId     string  `gorm:"column:variant_id;type:varchar(255)"`
	SKU           string  `gorm:"column:sku;type:varchar(100);not null"`
	Size          string  `gorm:"column:size;type:varchar(20)"`
	Discount      float32 `gorm:"column:discount;type:decimal(5,2)"`
	Price         float32 `gorm:"column:price;type:decimal(10,2)"`
	StockQuantity int     `gorm:"column:stock_quantity;type:int;not null"`
	IsAvailable   bool    `gorm:"column:is_available;type:bool"`

	ProductVariant ProductVariant `gorm:"foreignKey:variant_id;references:id"`
}

func (c *ProductSize) TableName() string {
	return "product_sizes"
}

func (c *ProductSize) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	if c.StockQuantity != 0 {
		c.IsAvailable = true
	}
	return
}
