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
	Config    *viper.Viper
	AdminRepo *adminrepo.AdminRepository
}

func NewAdminUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, config *viper.Viper, adminRepo *adminrepo.AdminRepository) *AdminUseCase {
	return &AdminUseCase{
		DB:        db,
		Log:       log,
		Validate:  validate,
		Config:    config,
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

func (c *AdminUseCase) Login(ctx context.Context, request *model.AdminLoginRequest) (*model.AdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	admin := new(entity.Admin)
	if err := c.AdminRepo.FindByEmail(tx, admin, request.Email); err != nil {
		c.Log.Warnf("Failed to query customer : %+v", err)

		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Invalid password for customer ID %s : %+v", admin.ID, err)

		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Generate refresh Token
	refreshToken, err := pkg.GenerateTokenAdmin(admin, c.Config.GetString("jwt.refreshToken"), c.Config.GetInt("jwt.expRefreshToken"))
	if err != nil {
		c.Log.Warnf("Failed generate refresh token: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	// Generate access Token
	accessToken, err := pkg.GenerateTokenAdmin(admin, c.Config.GetString("jwt.accessToken"), c.Config.GetInt("jwt.expAccessToken"))
	if err != nil {
		c.Log.Warnf("Failed generate refresh token: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	_, err = pkg.VerifyToken(refreshToken, c.Log, c.Config.GetString("jwt.refreshToken"))
	if err != nil {
		c.Log.Warnf("Failed generate refresh token: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	admin.RefreshToken = refreshToken

	// Store token to database
	if err := c.AdminRepo.Update(tx, admin); err != nil {
		c.Log.Warnf("Failed to save token to database: %+v", err)

		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.AdminToAuthResponse(admin, accessToken, refreshToken), nil
}
