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

func (c *CartItemRepository) FindByIdAndCustomerId(tx *gorm.DB, entity *entity.CartItem, id string, customerId string) error {
	return tx.
		Where("cart_items.id = ? AND cart_items.customer_id = ?", id, customerId).
		Joins("Product").
		Joins("Product.Category").
		Joins("Product.SubCategory").
		Preload("Product.ProductVariant").
		Preload("Product.ProductVariant.VariantImages").
		Preload("Product.ProductVariant.ProductSizes").
		Take(entity).Error
}

func (c *CartItemRepository) FindAllByCustomerId(tx *gorm.DB, customerId string) ([]*entity.CartItem, error) {
	var cartItems []*entity.CartItem

	err := tx.
		Where("cart_items.customer_id = ?", customerId).
		Joins("Product").
		Joins("Product.Category").
		Joins("Product.SubCategory").
		Preload("Product.ProductVariant").
		Preload("Product.ProductVariant.VariantImages").
		Preload("Product.ProductVariant.ProductSizes").
		Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (c *CartItemRepository) FindVariantId(tx *gorm.DB, entity *entity.CartItem, variantID string) error {
	return tx.Where("variant_id = ?", variantID).Take(entity).Error
}
