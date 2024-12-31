package http

import (
	"github.com/bwafi/trendora-backend/internal/model"
	adminusecase "github.com/bwafi/trendora-backend/internal/usecase/admin"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type AdminController struct {
	Log          *logrus.Logger
	AdminUseCase *adminusecase.AdminUseCase
}

func NewAdminUseCase(useCase *adminusecase.AdminUseCase, logger *logrus.Logger) *AdminController {
	return &AdminController{
		Log:          logger,
		AdminUseCase: useCase,
	}
}

func (c *AdminController) Register(ctx fiber.Ctx) error {
	request := new(model.CreateAdminRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.AdminResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.AdminUseCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.AdminResponse]{
		Data: response,
	})
}

func (c *AdminController) Login(ctx fiber.Ctx) error {
	request := new(model.AdminLoginRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.AdminResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.AdminUseCase.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.AdminResponse]{
		Data: response,
	})
}
