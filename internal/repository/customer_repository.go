package repository

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomersRepository struct {
	Repository[entity.Customers]
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomersRepository {
	return &CustomersRepository{
		Log: log,
	}
}

func (c *CustomersRepository) ExistsByEmail(tx *gorm.DB, email string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Customers{}).Where("email_address = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *CustomersRepository) ExistsByPhoneNumber(tx *gorm.DB, phoneNumber string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Customers{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
