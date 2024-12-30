package productusecase

import (
	"context"

	"github.com/bwafi/trendora-backend/internal/model"
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ProductRepository *productrepo.ProductRepository
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productRepo *productrepo.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		ProductRepository: productRepo,
	}
}

func (c *ProductUseCase) Create(ctx context.Context, request *model.CreateAddressRequest) (*model.ProductResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	return nil, nil
}
