package productusecase

import (
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	ProductRepository        *productrepo.ProductRepository
	CategoryRepository       *productrepo.CategoryRepository
	ProductImageRepository   *productrepo.ProductImageRepository
	ProductVariantRepository *productrepo.ProductVariantRepository
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productRepo *productrepo.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		ProductRepository: productRepo,
	}
}

// func (c *ProductUseCase) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
// 	tx := c.DB.WithContext(ctx).Begin()
// 	defer tx.Rollback()
//
// 	if err := c.Validate.Struct(request); err != nil {
// 		c.Log.Warnf("Invalid request body : %+v", err)
//
// 		message := pkg.ParseValidationErrors(err)
// 		return nil, fiber.NewError(fiber.StatusBadRequest, message)
// 	}
//
// 	category := new(entity.Category)
// 	if err := c.CategoryRepository.FindById(tx, category, request.CategoryId); err != nil {
// 		c.Log.Warnf("Category not found: %v", err)
// 		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
// 	}
//
// 	subCategory := new(entity.Category)
// 	if err := c.CategoryRepository.FindById(tx, subCategory, request.CategoryId); err != nil {
// 		c.Log.Warnf("Sub category not found: %v", err)
// 		return nil, fiber.NewError(fiber.StatusNotFound, "Sub category not found")
// 	}
// 	product := &entity.Product{
// 		StyleCode:      request.StyleCode,
// 		Name:           request.Name,
// 		Description:    request.Description,
// 		Gender:         request.Gender,
// 		CategoryId:     request.CategoryId,
// 		SubCategoryId:  request.SubCategoryId,
// 		BasePrice:      request.BasePrice,
// 		IsVisible:      request.IsVisible,
// 		ReleaseDate:    request.ReleaseDate,
// 		ProductImage:   []entity.ProductImage{},
// 		ProductVariant: []entity.ProductVariant{},
// 	}
//
// 	if err := c.ProductRepository.Create(tx, product); err != nil {
// 		c.Log.Warnf("Failed create product  to database : %+v", err)
// 		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
// 	}
//
// 	var productVariants []*entity.ProductVariant
// 	for i := range request.ProductVariant {
// 		productVariants = append(productVariants)
// 	}
//
// 	c.ProductVariantRepository.BulkCreate(tx, productVariants)
//
// 	return nil, nil
// }
