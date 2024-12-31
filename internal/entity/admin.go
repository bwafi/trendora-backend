package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID           string         `gorm:"column:id;type:uuid;primaryKey"`
	Name         string         `gorm:"column:name;type:varchar(50);not null"`
	Email        string         `gorm:"column:email;type:varchar(255);not null;index:unique"`
	Password     string         `gorm:"column:password;type:varchar(255);not null"`
	PhoneNumber  string         `gorm:"column:phone_number;type:varchar(20);unique"`
	RefreshToken string         `gorm:"column:refresh_token"`
	CreatedAt    int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (c *Admin) TableName() string {
	return "admins"
}

func (c *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
