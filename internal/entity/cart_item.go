package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID         string `gorm:"column:id;type:uuid;primaryKey"`
	CustomerId string `gorm:"column:customer_id;type:varchar(255);not null"`
	ProductId  string `gorm:"column:product_id;type:varchar(255);not null"`
	VariantId  string `gorm:"column:variant_id;type:varchar(255);not null"`
	Quantity   int    `gorm:"column:quantity;not null"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`

	Customers      *Customers      `gorm:"foreignKey:customer_id;references:id"`
	Product        *Product        `gorm:"foreignKey:product_id;references:id"`
	ProductVariant *ProductVariant `gorm:"foreignKey:variant_id;references:id"`
}

func (c *CartItem) TableName() string {
	return "cart_items"
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
