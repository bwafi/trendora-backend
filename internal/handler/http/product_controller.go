package http

import (
	"fmt"
	"strconv"

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

func (c *ProductController) Get(ctx fiber.Ctx) error {
	request := new(model.ProductGetRequest)
	productId := ctx.Params("productId")

	request.ID = productId

	response, err := c.ProductUseCase.Get(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.ProductResponse]{
		Data: response,
	})
}

func (c *ProductController) List(ctx fiber.Ctx) error {
	queries := ctx.Queries()

	page, err := strconv.Atoi(queries["page"])
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(queries["limit"])
	if err != nil || limit <= 0 {
		limit = 10
	}

	request := &model.ProductGetListRequest{
		ID:            "",
		Name:          queries["name"],
		Gender:        queries["gender"],
		CategoryId:    queries["category_id"],
		SubCategoryId: queries["sub_category_id"],
		Limit:         limit,
		Page:          page,
	}

	responses, pagination, err := c.ProductUseCase.List(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]*model.ProductResponse]{
		Data:   responses,
		Paging: pagination,
	})
}

func (c *ProductController) RecordView(ctx fiber.Ctx) error {
	request := new(model.ProductViewRequest)
	requestId := ctx.Params("productId")

	request.ID = requestId

	response, err := c.ProductUseCase.RecordView(ctx.RequestCtx(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.ProductResponse]{
		Data: response,
	})
}
