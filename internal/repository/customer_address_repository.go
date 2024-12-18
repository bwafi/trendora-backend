package repository

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/sirupsen/logrus"
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
