package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	ID           string `gorm:"column:id;type:uuid;primaryKey"`
	ProductId    string `gorm:"column:product_id;type:varchar(255)"`
	VarianId     string `gorm:"column:variant_id;type:varchar(255)"`
	ImageUrl     string `gorm:"column:image_url;type:varchar(255)"`
	DisplayOrder int    `gorm:"column:display_order;type:int"`
}

func (c *ProductImage) TableName() string {
	return "proucts"
}

func (c *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
