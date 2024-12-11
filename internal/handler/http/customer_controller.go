package http

import (
	"time"

	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CustomerController struct {
	Log          *logrus.Logger
	CustomerCase *usecase.CustomerUseCase
	Config       *viper.Viper
}

func NewCustomerController(useCase *usecase.CustomerUseCase, logger *logrus.Logger, viper *viper.Viper) *CustomerController {
	return &CustomerController{
		Log:          logger,
		CustomerCase: useCase,
		Config:       viper,
	}
}

func (c *CustomerController) Register(ctx fiber.Ctx) error {
	request := new(model.CustomerRegisterRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Invalid request body",
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register customer : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Customer registration failed",
			Errors: &model.ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CustomerResponse]{
		Status:  "Success",
		Message: "Customer registration successful",
		Data:    response,
	})
}

func (c *CustomerController) Login(ctx fiber.Ctx) error {
	request := new(model.CustomerLoginRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Invalid request body",
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register customer : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Customer registration failed",
			Errors: &model.ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			},
		})
	}

	timeDuration := c.Config.GetInt("jwt.expRefreshToken")

	ctx.Cookie(&fiber.Cookie{
		Name:     "_SID_Trendora",
		Value:    *response.Token,
		Expires:  time.Now().Add(time.Duration(timeDuration)),
		HTTPOnly: true,
		Secure:   true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CustomerResponse]{
		Status:  "Success",
		Message: "Customer registration successful",
		Data:    response,
	})
}

func (c *CustomerController) Update(ctx fiber.Ctx) error {
	request := new(model.CustomerUpdateRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Invalid request body",
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update customer : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Customer update failed",
			Errors: &model.ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CustomerResponse]{
		Status:  "Success",
		Message: "Customer update successful",
		Data:    response,
	})
}

func (c *CustomerController) Delete(ctx fiber.Ctx) error {
	request := new(model.CustomerDeleteRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Invalid request body",
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to delete customer : %+v", err)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(model.WebResponse[*model.CustomerResponse]{
			Status:  "Failed",
			Message: "Customer update failed",
			Errors: &model.ErrorResponse{
				Code:    statusCode,
				Message: err.Error(),
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CustomerResponse]{
		Status:  "Success",
		Message: "Customer delete successful",
		Data:    response,
	})
}
