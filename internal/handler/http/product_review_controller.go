package http

import (
	"github.com/bwafi/trendora-backend/internal/handler/middleware"
	"github.com/bwafi/trendora-backend/internal/model"
	productusecase "github.com/bwafi/trendora-backend/internal/usecase/product"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ProductReviewController struct {
	Log                  *logrus.Logger
	ProductReviewUseCase *productusecase.ProductReviewUsecase
}

func NewProductReviewController(useCase *productusecase.ProductReviewUsecase, log *logrus.Logger) *ProductReviewController {
	return &ProductReviewController{
		Log:                  log,
		ProductReviewUseCase: useCase,
	}
}

func (c *ProductReviewController) Create(ctx fiber.Ctx) error {
	request := new(model.ProductReviewRequest)
	auth := middleware.GetUser(ctx)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.ProductReviewResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	request.CustomerID = auth.ID

	response, err := c.ProductReviewUseCase.Create(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.ProductReviewResponse]{
		Data: response,
	})
}
