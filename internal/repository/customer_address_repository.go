package repository

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomersAddressRepository struct {
	Repository[entity.CustomerAddresses]
	Log *logrus.Logger
}

func NewCustomerAddressesRepository(log *logrus.Logger) *CustomersAddressRepository {
	return &CustomersAddressRepository{
		Log: log,
	}
}

func (c *CustomersAddressRepository) FindByIdAndCustomerId(tx *gorm.DB, entity *entity.CustomerAddresses, id string, customerId string) error {
	return tx.Where("id = ? AND customer_id = ? ", id, customerId).Take(entity).Error
}
