package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID            string  `gorm:"column:id;type:uuid;primaryKey"`
	ProductId     string  `gorm:"column:product_id;type:varchar(255)"`
	SKU           string  `gorm:"column:sku;type:varchar(100);not null"`
	ColorName     string  `gorm:"column:color_name;type:varchar(50)"`
	Size          string  `gorm:"column:size;type:varchar(20)"`
	Discount      float32 `gorm:"column:discount;type:decimal(5,2)"`
	Price         float32 `gorm:"column:price;type:decimal(10,2)"`
	StockQuantity int     `gorm:"column:stock_quantity;type:int;not null"`
	Weight        float32 `gorm:"column:weight;type:decimal(8,2)"`
	IsAvailable   bool    `gorm:"column:is_available;type:bool"`

	VariantImage []VariantImage `gorm:"foreignKey:variant_id;references:id"`
	Product      Product        `gorm:"foreignKey:product_id;references:id"`
}

func (c *ProductVariant) TableName() string {
	return "prouct_variants"
}

func (c *ProductVariant) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
