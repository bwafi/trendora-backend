package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerAddresses struct {
	ID            string         `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	CustomerID    string         `gorm:"column:customer_id;type:uuid;not null" json:"customer_id"`
	RecipientName string         `gorm:"column:recipient_name;type:varchar(100);not null" json:"recipient_name"`
	PhoneNumber   string         `gorm:"column:phone_number;type:varchar(20)" json:"phone_number"`
	AddressType   string         `gorm:"column:address_type;type:varchar(10);not null" json:"address_type"`
	City          string         `gorm:"column:city;type:varchar(100);not null" json:"city"`
	Province      string         `gorm:"column:province;type:varchar(100);not null" json:"province"`
	SubDistrict   string         `gorm:"column:sub_district;type:varchar(100)" json:"sub_district"`
	PostalCode    string         `gorm:"column:postal_code;type:varchar(10)" json:"postal_code"`
	CourierNotes  string         `gorm:"column:courier_notes;type:text" json:"courier_notes"`
	CreatedAt     int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Customers     Customers      `gorm:"foreignKey:customer_id;references:id"`
}

func (c *CustomerAddresses) TableName() string {
	return "customer_addresses"
}

func (c *CustomerAddresses) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
