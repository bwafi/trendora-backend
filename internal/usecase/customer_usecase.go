package usecase

import (
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	Validate            *validator.Validate
	CustomersRepository *repository.CustomersRepository
}

func NewCustomerUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, customersRepository *repository.CustomersRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                  db,
		Log:                 log,
		Validate:            validate,
		CustomersRepository: customersRepository,
	}
}
