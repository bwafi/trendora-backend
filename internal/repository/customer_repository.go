package repository

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/sirupsen/logrus"
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
