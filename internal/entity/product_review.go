package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductReview struct {
	ID         string         `gorm:"column:id;type:uuid;primaryKey"`
	ProductId  string         `gorm:"column:product_id;type:varchar(255)"`
	CustomerID string         `gorm:"column:customer_id;type:uuid;not null" json:"customer_id"`
	Rating     float64        `gorm:"column:rating;type:decimal(2,1)"`
	Comment    string         `gorm:"column:comment;type:TEXT"`
	CreatedAt  int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Customer   Customers      `gorm:"foreignKey:customer_id;references:id"`
	Product    Product        `gorm:"foreignKey:product_id;references:id"`
}

func (c *ProductReview) TableName() string {
	return "product_reviews"
}

func (c *ProductReview) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
