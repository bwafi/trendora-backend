package customerrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomersAddressRepository struct {
	repository.Repository[entity.CustomerAddresses]
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

func (c *CustomersAddressRepository) FindAllByCustomerId(tx *gorm.DB, customerId string) ([]entity.CustomerAddresses, error) {
	var addresses []entity.CustomerAddresses

	if err := tx.Where("customer_id = ?", customerId).Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}
