package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            string         `gorm:"column:id;type:uuid;primaryKey"`
	StyleCode     string         `gorm:"column:style_code;type:varchar(50);not null"`
	Name          string         `gorm:"column:name;type:varchar(255);not null"`
	Description   string         `gorm:"column:description;type:text"`
	Gender        string         `gorm:"column:gender;type:varchar(10)"`
	CategoryId    string         `gorm:"column:category_id;type:varchar(255)"`
	SubCategoryId string         `gorm:"column:sub_category_id;type:varchar(255)"`
	BasePrice     float32        `gorm:"column:base_price;type:decimal(10,2)"`
	IsVisible     bool           `gorm:"column:is_visible;type:bool;default:true"`
	ReleaseDate   int64          `gorm:"column:release_date;type:bigint"`
	CreatedAt     int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (c *Product) TableName() string {
	return "proucts"
}

func (c *Product) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}