package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
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

func (c *ProductRepository) FindAllProducts(tx *gorm.DB, request *model.ProductGetListRequest) ([]*entity.Product, error) {
	var products []*entity.Product

	query := tx.
		Joins("Category").
		Joins("SubCategory").
		Preload("ProductVariant").
		Preload("ProductVariant.VariantImages").
		Preload("ProductVariant.ProductSizes")

	if request.CategoryId != "" {
		query = query.Where("products.category_id = ?", request.CategoryId)
	}
	if request.SubCategoryId != "" {
		query = query.Where("products.sub_category_id = ?", request.SubCategoryId)
	}
	if request.Name != "" {
		query = query.Where("products.name ILIKE ?", "%"+request.Name+"%")
	}

	if request.Gender != "" {
		query = query.Where("products.gender = ?", request.Gender)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
