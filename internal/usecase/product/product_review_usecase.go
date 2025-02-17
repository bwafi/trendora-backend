package productusecase

import (
	"context"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/model/converter"
	customerrepo "github.com/bwafi/trendora-backend/internal/repository/customer"
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductReviewUsecase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	Validate                *validator.Validate
	ProductReviewRepository *productrepo.ProductReviewRepository
	ProductRepository       *productrepo.ProductRepository
	CustomerRepository      *customerrepo.CustomerRepository
}

func NewProductReviewUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productReviewRepo *productrepo.ProductReviewRepository, productRepo *productrepo.ProductRepository, customerRepo *customerrepo.CustomerRepository) *ProductReviewUsecase {
	return &ProductReviewUsecase{
		DB:                      db,
		Log:                     log,
		Validate:                validate,
		ProductReviewRepository: productReviewRepo,
		ProductRepository:       productRepo,
		CustomerRepository:      customerRepo,
	}
}

func (c *ProductReviewUsecase) Create(ctx context.Context, request *model.ProductReviewRequest) (*model.ProductReviewResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product := new(entity.Product)
	if err := c.ProductRepository.FindById(tx, product, request.ProductId); err != nil {
		c.Log.Warnf("Product with id %s not found", request.ProductId)
		return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	customer := new(entity.Customers)
	if err := c.CustomerRepository.FindById(tx, customer, &request.CustomerID); err != nil {
		c.Log.Warnf("Customer with id %s not found", request.ProductId)
		return nil, fiber.NewError(fiber.StatusNotFound, "Customer not found")
	}

	productReview := &entity.ProductReview{
		ProductId:  request.ProductId,
		CustomerID: request.CustomerID,
		Rating:     request.Rating,
		Comment:    request.Comment,
	}

	if err := c.ProductReviewRepository.Create(tx, productReview); err != nil {
		c.Log.Warnf("Failed save product review : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.ProductReviewToResponse(productReview), nil
}
