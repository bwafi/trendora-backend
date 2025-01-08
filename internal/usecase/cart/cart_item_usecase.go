package cartusecase

import (
	cartrepo "github.com/bwafi/trendora-backend/internal/repository/cart"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartItemUseCase struct {
	DB           *gorm.DB
	Log          *logrus.Logger
	Validate     *validator.Validate
	CartItemRepo *cartrepo.CartItemRepository
}

func NewCartItemUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, cartItemRepo *cartrepo.CartItemRepository) *CartItemUseCase {
	return &CartItemUseCase{
		DB:           db,
		Log:          log,
		Validate:     validate,
		CartItemRepo: cartItemRepo,
	}
}

// func (c *CartItemUseCase) Create(ctx context.Context, request)  {
//
// }
