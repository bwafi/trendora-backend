package http

import (
	"fmt"

	"github.com/bwafi/trendora-backend/internal/model"
	productusecase "github.com/bwafi/trendora-backend/internal/usecase/product"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Log            *logrus.Logger
	ProductUseCase *productusecase.ProductUseCase
}

func NewProductController(useCase *productusecase.ProductUseCase, logger *logrus.Logger) *ProductController {
	return &ProductController{
		Log:            logger,
		ProductUseCase: useCase,
	}
}

func (c *ProductController) Create(ctx fiber.Ctx) error {
	request := new(model.CreateProductRequest)

	ctx.Request().Header.Add(fiber.HeaderContentType, fiber.MIMEMultipartForm)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.CustomerResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	for i := range request.ProductImages {
		file, err := ctx.FormFile(fmt.Sprintf("product_images-%d", i))
		if err != nil {
			continue
		}
		request.ProductImages[i].Image = file
	}

	for i, variant := range request.ProductVariants {
		for j := range variant.VariantImages {
			file, err := ctx.FormFile(fmt.Sprintf("variant-%d-image-%d", i, j))
			if err != nil {
				continue
			}
			request.ProductVariants[i].VariantImages[j].Image = file
		}
	}

	response, err := c.ProductUseCase.Create(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.ProductResponse]{
		Data: response,
	})
}
