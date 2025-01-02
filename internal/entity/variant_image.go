package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VariantImage struct {
	ID           string `gorm:"column:id;type:uuid;primaryKey"`
	VarianId     string `gorm:"column:variant_id;type:varchar(255)"`
	ImageUrl     string `gorm:"column:image_url;type:varchar(255)"`
	DisplayOrder int    `gorm:"column:display_order;type:int"`

	ProductVariant ProductVariant `gorm:"foreignKey:variant_id;references:id"`
}

func (c *VariantImage) TableName() string {
	return "variant_images"
}

func (c *VariantImage) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
