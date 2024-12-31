package adminusecase

import (
	"context"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/model/converter"
	adminrepo "github.com/bwafi/trendora-backend/internal/repository/admin"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminUseCase struct {
	DB        *gorm.DB
	Log       *logrus.Logger
	Validate  *validator.Validate
	AdminRepo *adminrepo.AdminRepository
}

func NewAdminUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, config *viper.Viper, adminRepo *adminrepo.AdminRepository) *AdminUseCase {
	return &AdminUseCase{
		DB:        db,
		Log:       log,
		Validate:  validate,
		AdminRepo: adminRepo,
	}
}

func (c *AdminUseCase) Create(ctx context.Context, request *model.CreateAdminRequest) (*model.AdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	exist, err := c.AdminRepo.ExistsByEmail(tx, &request.Email)
	if err != nil {
		c.Log.Warnf("Failed to check email existence : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	if exist > 0 {
		c.Log.Warnf("Email already in use : %+v", err)
		return nil, fiber.NewError(fiber.StatusConflict, "Email already in use")
	}

	exist, err = c.AdminRepo.ExistsByPhoneNumber(tx, &request.PhoneNumber)
	if err != nil {
		c.Log.Warnf("Failed to check phone number existence : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	if exist > 0 {
		c.Log.Warnf("Phone number already in use : %+v", err)
		return nil, fiber.NewError(fiber.StatusConflict, "Phone number already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed encrypt password : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	admin := &entity.Admin{
		Name:        request.Name,
		Email:       request.Email,
		Password:    string(hashedPassword),
		PhoneNumber: request.PhoneNumber,
	}

	if err := c.AdminRepo.Create(tx, admin); err != nil {
		c.Log.Warnf("Failed create customer to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.AdminToAdminResponse(admin), nil
}
