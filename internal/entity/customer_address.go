package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerAddresses struct {
	ID        string         `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	UserId    string         `gorm:"column:user_id"`
	CreatedAt int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Customers Customers      `gorm:"foreignKey:user_id;references:id"`
}

func (c *CustomerAddresses) TableName() string {
	return "customer_addresses"
}

func (c *CustomerAddresses) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
