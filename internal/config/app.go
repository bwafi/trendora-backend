package config

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/bwafi/trendora-backend/internal/handler/http/route"
	"github.com/bwafi/trendora-backend/internal/handler/middleware"
	customerrepo "github.com/bwafi/trendora-backend/internal/repository/customer"
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

	customerUseCase := customerusecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, config.Config, customerRepository, customerSessionRepository)
	customerAddressUsecase := customerusecase.NewCustomerAddressUsecase(config.DB, config.Log, config.Validate, config.Config, customerAddressRepository)

	customerController := http.NewCustomerController(customerUseCase, config.Log)
	cusomerAddressController := http.NewCustomerAddressController(customerAddressUsecase, config.Log, config.Config)

	authMiddleware := middleware.AuthMiddleware(customerUseCase)

	routeConfig := route.RouteConfig{
		App:                       config.App,
		CustomerController:        customerController,
		CustomerAddressController: cusomerAddressController,
		AuthMiddleware:            authMiddleware,
	}

	routeConfig.Setup()
}
