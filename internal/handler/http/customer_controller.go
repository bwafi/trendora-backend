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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CustomerResponse]{
		Data: response,
	})
}

func (c *CustomerController) Login(ctx fiber.Ctx) error {
	request := new(model.CustomerLoginRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CustomerResponse]{
		Data: response,
	})
}

func (c *CustomerController) Update(ctx fiber.Ctx) error {
	request := new(model.CustomerUpdateRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Update(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CustomerResponse]{
		Data: response,
	})
}

func (c *CustomerController) Delete(ctx fiber.Ctx) error {
	request := new(model.CustomerDeleteRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerCase.Delete(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CustomerResponse]{
		Data: response,
	})
}
