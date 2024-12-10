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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	Validate            *validator.Validate
	CustomersRepository *repository.CustomersRepository
	Config              *viper.Viper
}

func NewCustomerUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, config *viper.Viper, customersRepository *repository.CustomersRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                  db,
		Log:                 log,
		Validate:            validate,
		CustomersRepository: customersRepository,
		Config:              config,
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

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	if request.EmailAddress != nil {
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

	if request.PhoneNumber != nil {
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
		c.Log.Warnf("Failed encrypt password : %+v", err)
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
		c.Log.Warnf("Failed create customer to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Login(ctx context.Context, request *model.CustomerLoginRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)

		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	customer := new(entity.Customers)
	err := c.CustomersRepository.FindByEmailOrPhone(tx, customer, request.EmailAddress, request.PhoneNumber)
	if err != nil {
		c.Log.Warnf("Failed to query customer : %+v", err)

		if request.EmailAddress != nil {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
		}

		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid phone or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		c.Log.Warnf("Invalid password for customer ID %s : %+v", customer.ID, err)

		if request.EmailAddress != nil {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
		}

		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid phone or password")
	}

	// Generate Token
	Token, errGenerate := pkg.GenerateRefreshToken(customer, c.Config)
	if errGenerate != nil {
		return nil, errGenerate
	}

	customer.Token = &Token
	if err := c.CustomersRepository.Update(tx, customer); err != nil {
		c.Log.Warnf("Failed to save token to database: %+v", err)

		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Update(ctx context.Context, request *model.CustomerUpdateRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)

		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	var password string
	if request.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.Warnf("Failed encrypt password : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		password = string(hashPassword)
	}

	customer := &entity.Customers{
		ID:           request.ID,
		Name:         request.Name,
		EmailAddress: request.EmailAddress,
		PhoneNumber:  request.PhoneNumber,
		Password:     string(password),
		DateOfBirth:  request.DateOfBirth,
		Gender:       request.Gender,
	}

	if err := c.CustomersRepository.Update(tx, customer); err != nil {
		c.Log.Warnf("Failed update customer to database : %+v", err)

		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Delete(ctx context.Context, request *model.CustomerDeleteRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)

		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	customer := new(entity.Customers)
	err := c.CustomersRepository.FindById(tx, customer, &request.ID)
	if err != nil {
		c.Log.Warnf("Failed to find customer : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "Customer not found")
	}

	if err := c.CustomersRepository.Delete(tx, customer); err != nil {
		c.Log.Warnf("Failed to delete customer : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CustomerToResponse(customer), nil
}
