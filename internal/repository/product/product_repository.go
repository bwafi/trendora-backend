package productrepo

import (
	"fmt"
	"math"

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

func (c *ProductRepository) FindAllProducts(tx *gorm.DB, request *model.ProductGetListRequest) ([]*entity.Product, *model.PageMetadata, error) {
	var products []*entity.Product
	var totalItems int64

	query := tx.
		Debug().
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

	err := query.Model(&entity.Product{}).Count(&totalItems).Error
	fmt.Println("Error counting total items:", err)
	if err != nil {
		fmt.Println("Error counting total items:", err)
		return nil, nil, err
	}

	pageSize := 10
	if request.Limit > 0 {
		pageSize = request.Limit
	}

	currentPage := 1
	if request.Page > 0 {
		currentPage = request.Page
	}

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pageSize)))

	offset := (currentPage - 1) * pageSize
	err = query.
		Limit(pageSize).
		Offset(offset).
		Find(&products).Error
	if err != nil {
		return nil, nil, err
	}

	pagination := &model.PageMetadata{
		Page:       pageSize,
		Size:       pageSize,
		TotalItem:  totalItems,
		TotalPages: totalPages,
	}

	return products, pagination, nil
}
