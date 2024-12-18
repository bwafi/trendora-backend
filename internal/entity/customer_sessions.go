package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerSessions struct {
	ID           string    `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	CustomerID   string    `gorm:"column:customer_id;not null" json:"customer_id"`
	RefreshToken string    `gorm:"column:refresh_token;not null"`
	IsRevoked    bool      `gorm:"column:is_revoked;default:false"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	CreatedAt    int64     `gorm:"column:created_at;autoCreateTime:milli"`
}

func (c *CustomerSessions) TableName() string {
	return "customer_sessions"
}

// TODO: Make sure without method BeforeCreaate because using default uuid_genereate_v4()
func (c *CustomerSessions) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
