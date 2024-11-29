package entity

import "gorm.io/gorm"

type Customers struct {
	ID           string         `gorm:"primaryKey;column:id"`
	EmailAddress string         `gorm:"column:email_address"`
	PhoneNumber  string         `gorm:"column:phone_number"`
	Password     string         `gorm:"column:password"`
	CreatedAt    int64          `gorm:"column:craeted_at;autoCreateTime:milli;"`
	UpdatedAt    int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (c *Customers) TableName() string {
	return "customers"
}
