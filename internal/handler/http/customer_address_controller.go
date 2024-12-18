package http

import (
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CustomerAddressController struct {
	Log                 *logrus.Logger
	CustomerAddressCase *usecase.CustomerAddressUsecase
	Config              *viper.Viper
}

func NewCustomerAddressController(useCase *usecase.CustomerAddressUsecase, logger *logrus.Logger, viper *viper.Viper) *CustomerAddressController {
	return &CustomerAddressController{
		Log:                 logger,
		CustomerAddressCase: useCase,
		Config:              viper,
	}
}

func (c *CustomerAddressController) Create(ctx fiber.Ctx) error {
	request := new(model.CreateAddressRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.AddressResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerAddressCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register customer : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(model.WebResponse[*model.AddressResponse]{
			Errors: &model.ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.AddressResponse]{
		Data: response,
	})
}