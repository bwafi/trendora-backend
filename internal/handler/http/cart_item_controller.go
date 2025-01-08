package http

import (
	"github.com/bwafi/trendora-backend/internal/model"
	cartusecase "github.com/bwafi/trendora-backend/internal/usecase/cart"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type CartItemController struct {
	Log             *logrus.Logger
	CartItemUseCase *cartusecase.CartItemUseCase
}

func NewCartItemController(log *logrus.Logger, CartItemUseCase *cartusecase.CartItemUseCase) *CartItemController {
	return &CartItemController{
		Log:             log,
		CartItemUseCase: CartItemUseCase,
	}
}

func (c *CartItemController) Create(ctx fiber.Ctx) error {
	request := new(model.CartItemRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CartItemUseCase.Create(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CartItemResponse]{
		Data: response,
	})
}
