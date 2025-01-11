package cartusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/model/converter"
	cartrepo "github.com/bwafi/trendora-backend/internal/repository/cart"
	customerrepo "github.com/bwafi/trendora-backend/internal/repository/customer"
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartItemUseCase struct {
	DB           *gorm.DB
	Log          *logrus.Logger
	Validate     *validator.Validate
	CartItemRepo *cartrepo.CartItemRepository
	CustomerRepo *customerrepo.CustomerRepository
	ProductRepo  *productrepo.ProductRepository
	VariantRepo  *productrepo.ProductVariantRepository
}

func NewCartItemUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, cartItemRepo *cartrepo.CartItemRepository, CustomerRepo *customerrepo.CustomerRepository, ProductRepo *productrepo.ProductRepository, VariantRepo *productrepo.ProductVariantRepository,
) *CartItemUseCase {
	return &CartItemUseCase{
		DB:           db,
		Log:          log,
		Validate:     validate,
		CartItemRepo: cartItemRepo,
		CustomerRepo: CustomerRepo,
		ProductRepo:  ProductRepo,
		VariantRepo:  VariantRepo,
	}
}

func (c *CartItemUseCase) Create(ctx context.Context, request *model.CartItemRequest) (*model.CartItemResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	customer := new(entity.Customers)
	if err := c.CustomerRepo.FindById(tx, customer, &request.CustomerId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Customer with id %s not found", request.CustomerId)
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Customer with id %s not found", request.CustomerId))
		}

		c.Log.Warnf("Error : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	product := new(entity.Product)
	if err := c.ProductRepo.FindById(tx, product, request.ProductId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Product with id %s not found", request.CustomerId)
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Product with id %s not found", request.CustomerId))
		}

		c.Log.Warnf("Error : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	ProductVariant := new(entity.ProductVariant)
	if err := c.VariantRepo.FindById(tx, ProductVariant, request.VariantId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Product variant with id %s not found", request.CustomerId)
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Product variant with id %s not found", request.CustomerId))
		}

		c.Log.Warnf("Error : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	// Check if exist increase quantity only
	cartItem := new(entity.CartItem)
	err := c.CartItemRepo.FindVariantId(tx, cartItem, request.VariantId)

	if err == nil {
		cartItem.Quantity += request.Quantity

		if err := c.CartItemRepo.Update(tx, cartItem); err != nil {
			c.Log.Warnf("Failed to add quantity to cart : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		// Check if err record not found
	} else if errors.Is(err, gorm.ErrRecordNotFound) {

		cartItem = &entity.CartItem{
			CustomerId: request.CustomerId,
			ProductId:  request.ProductId,
			VariantId:  request.VariantId,
			Quantity:   request.Quantity,
		}

		if err := c.CartItemRepo.Create(tx, cartItem); err != nil {
			c.Log.Warnf("Failed to save to cart : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		// if err not record not found
	} else {
		c.Log.Warnf("Failed to find cart item: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CartItemToReponse(cartItem), nil
}

func (c *CartItemUseCase) Update(ctx context.Context, request *model.CartItemUpdateRequest) (*model.CartItemResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	cartItem := new(entity.CartItem)
	if err := c.CartItemRepo.FindById(tx, cartItem, request.ID); err != nil {
		c.Log.Warnf("Item with id %s not found", request.ID)
		return nil, fiber.NewError(fiber.StatusNotFound, "Cart Product not found")
	}

	cartItem.Quantity += request.Quantity

	if err := c.CartItemRepo.Update(tx, cartItem); err != nil {
		c.Log.Warnf("Failed Update Item with id %s", request.ID)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.CartItemToReponse(cartItem), nil
}
