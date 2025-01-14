package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	repository.Repository[entity.Product]
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (c *ProductRepository) FindDetailsProduct(tx *gorm.DB, product *entity.Product, id string) error {
	return tx.
		Where("products.id = ?", id).
		Joins("Category").
		Joins("SubCategory").
		Preload("ProductVariant").
		Preload("ProductVariant.VariantImages").
		Preload("ProductVariant.ProductSizes").
		Take(product).Error
}
