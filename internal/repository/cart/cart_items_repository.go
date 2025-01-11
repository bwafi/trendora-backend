package cartrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (c *CartItemRepository) FindById(tx *gorm.DB, entity *entity.CartItem, id string) error {
	return tx.Where("id = ?", id).Take(entity).Error
}

func (c *CartItemRepository) FindVariantId(tx *gorm.DB, entity *entity.CartItem, variantID string) error {
	return tx.Where("variant_id = ?", variantID).Take(entity).Error
}
