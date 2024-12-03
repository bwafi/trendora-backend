package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customers struct {
	ID           string         `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	EmailAddress string         `gorm:"column:email_address"`
	PhoneNumber  string         `gorm:"column:phone_number"`
	Name         string         `gorm:"column:name;not null"`
	Gender       bool           `gorm:"column:gender"`
	DateOfBirth  int64          `gorm:"column:date_of_birth"`
	Password     string         `gorm:"column:password;not null"`
	Token        string         `gorm:"column:token"`
	CreatedAt    int64          `gorm:"column:craeted_at;autoCreateTime:milli;"`
	UpdatedAt    int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (c *Customers) TableName() string {
	return "customers"
}

func (c *Customers) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
