package cartrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type CartItemRepository struct {
	repository.Repository[entity.CartItem]
	Log *logrus.Logger
}

func NewCartItemRepository(log *logrus.Logger) *CartItemRepository {
	return &CartItemRepository{
		Log: log,
	}
}
