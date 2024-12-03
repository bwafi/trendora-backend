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
	"golang.org/x/crypto/bcrypt"
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

func (c *CustomerUseCase) Create(ctx context.Context, request *model.CustomerRegisterRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// register validation phone and email
	c.Validate.RegisterStructValidation(pkg.MustValidRegisterCustomer, model.CustomerRegisterRequest{})
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if request.EmailAddress != "" {
		exists, err := c.CustomersRepository.ExistsByEmail(tx, request.EmailAddress)
		if err != nil {
			c.Log.Warnf("Failed to check email existence : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		if exists > 0 {
			c.Log.Warnf("Email already in use : %+v", err)
			return nil, fiber.NewError(fiber.StatusBadRequest, "Email already in use")
		}
	}

	if request.PhoneNumber != "" {
		exists, err := c.CustomersRepository.ExistsByPhoneNumber(tx, request.PhoneNumber)
		if err != nil {
			c.Log.Warnf("Failed to check phone number existence : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		if exists > 0 {
			c.Log.Warnf("Phone number already in use : %+v", err)
			return nil, fiber.NewError(fiber.StatusBadRequest, "Phone number already in use")
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	customer := &entity.Customers{
		Name:         request.Name,
		EmailAddress: request.EmailAddress,
		PhoneNumber:  request.PhoneNumber,
		Password:     string(password),
		DateOfBirth:  request.DateOfBirth,
		Gender:       request.Gender,
	}

	if err := c.CustomersRepository.Create(tx, customer); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CustomerToResposne(customer), nil
}
