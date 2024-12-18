package usecase

import (
	"context"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/model/converter"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CustomerAddressUsecase struct {
	DB                        *gorm.DB
	Validate                  *validator.Validate
	Config                    *viper.Viper
	Log                       *logrus.Logger
	CustomerAddressRepository *repository.CustomersAddressRepository
}

func NewCustomerAddressUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, config *viper.Viper, customersAddressRepository *repository.CustomersAddressRepository) *CustomerAddressUsecase {
	return &CustomerAddressUsecase{
		DB:                        db,
		Validate:                  validate,
		Config:                    config,
		Log:                       log,
		CustomerAddressRepository: &repository.CustomersAddressRepository{},
	}
}

func (c *CustomerAddressUsecase) Create(ctx context.Context, request *model.CreateAddressRequest) (*model.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)

		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	customerAddress := &entity.CustomerAddresses{
		CustomerID:    request.CustomerID,
		RecipientName: request.RecipientName,
		PhoneNumber:   request.PhoneNumber,
		AddressType:   request.AddressType,
		City:          request.City,
		Province:      request.Province,
		SubDistrict:   request.SubDistrict,
		PostalCode:    request.PostalCode,
		CourierNotes:  request.CourierNotes,
	}

	if err := c.CustomerAddressRepository.Create(tx, customerAddress); err != nil {
		c.Log.Warnf("Failed create customer to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CustomerAddressToResponse(customerAddress), nil
}
