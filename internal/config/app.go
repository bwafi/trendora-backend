package config

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/bwafi/trendora-backend/internal/handler/http/route"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/bwafi/trendora-backend/internal/usecase"
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
	customerRepository := repository.NewCustomerRepository(config.Log)

	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, config.Config, customerRepository)

	customerController := http.NewCustomerController(customerUseCase, config.Log, config.Config)

	routeConfig := route.RouteConfig{
		App:                config.App,
		CustomerController: customerController,
	}

	routeConfig.Setup()
}
