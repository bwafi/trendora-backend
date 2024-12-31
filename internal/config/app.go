package config

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/bwafi/trendora-backend/internal/handler/http/route"
	"github.com/bwafi/trendora-backend/internal/handler/middleware"
	adminrepo "github.com/bwafi/trendora-backend/internal/repository/admin"
	customerrepo "github.com/bwafi/trendora-backend/internal/repository/customer"
	adminusecase "github.com/bwafi/trendora-backend/internal/usecase/admin"
	customerusecase "github.com/bwafi/trendora-backend/internal/usecase/customer"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	customerRepository := customerrepo.NewCustomerRepository(config.Log)
	customerSessionRepository := customerrepo.NewCustomerSessionRepository(config.Log)
	customerAddressRepository := customerrepo.NewCustomerAddressesRepository(config.Log)
	adminRepository := adminrepo.NewAdminrRepository(config.Log)

	customerUseCase := customerusecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, config.Config, customerRepository, customerSessionRepository)
	customerAddressUsecase := customerusecase.NewCustomerAddressUsecase(config.DB, config.Log, config.Validate, config.Config, customerAddressRepository)
	adminUseCase := adminusecase.NewAdminUseCase(config.DB, config.Log, config.Validate, config.Config, adminRepository)

	customerController := http.NewCustomerController(customerUseCase, config.Log)
	cusomerAddressController := http.NewCustomerAddressController(customerAddressUsecase, config.Log, config.Config)
	adminController := http.NewAdminUseCase(adminUseCase, config.Log)

	customerAuthMiddleware := middleware.CustomerAuthMiddleware(customerUseCase)
	adminAuthMiddleware := middleware.AdminAuthMiddleware(adminUseCase)

	routeConfig := route.RouteConfig{
		App:                       config.App,
		CustomerController:        customerController,
		CustomerAddressController: cusomerAddressController,
		AdminController:           adminController,
		CustomerAuthMiddleware:    customerAuthMiddleware,
		AdminAuthMiddleware:       adminAuthMiddleware,
	}

	routeConfig.Setup()
}
