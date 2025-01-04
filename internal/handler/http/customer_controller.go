package http

import (
	"github.com/bwafi/trendora-backend/internal/model"
	customerusecase "github.com/bwafi/trendora-backend/internal/usecase/customer"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	Log          *logrus.Logger
	CustomerCase *customerusecase.CustomerUseCase
}

func NewCustomerController(useCase *customerusecase.CustomerUseCase, logger *logrus.Logger) *CustomerController {
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

	response, err := c.CustomerCase.Create(ctx.RequestCtx(), request)
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

	response, err := c.CustomerCase.Login(ctx.RequestCtx(), request)
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

	response, err := c.CustomerCase.Update(ctx.RequestCtx(), request)
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

	response, err := c.CustomerCase.Delete(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CustomerResponse]{
		Data: response,
	})
}
