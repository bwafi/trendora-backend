package http

import (
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	Log          *logrus.Logger
	CustomerCase *usecase.CustomerUseCase
}

func NewCustomerController(useCase *usecase.CustomerUseCase, logger *logrus.Logger) *CustomerController {
	return &CustomerController{
		Log:          logger,
		CustomerCase: useCase,
	}
}

func (c *CustomerController) Register(ctx fiber.Ctx) error {
	request := new(model.CustomerRegisterRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{Errors: &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		}})
	}

	response, err := c.CustomerCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Errors: &model.ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}})
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}