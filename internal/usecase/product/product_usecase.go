package productusecase

import (
	"context"
	"fmt"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/model/converter"
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	Cloudinary               *cloudinary.Cloudinary
	ProductRepository        *productrepo.ProductRepository
	CategoryRepository       *productrepo.CategoryRepository
	ProductImageRepository   *productrepo.ProductImageRepository
	VariantImageRepository   *productrepo.VariantImageRepository
	ProductVariantRepository *productrepo.ProductVariantRepository
	ProductSizeRepository    *productrepo.ProductSizeRepository
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, cloudinary *cloudinary.Cloudinary, productRepo *productrepo.ProductRepository, categoryRepository *productrepo.CategoryRepository, productImageRepo *productrepo.ProductImageRepository, variantImageRepo *productrepo.VariantImageRepository, productVariantRepo *productrepo.ProductVariantRepository, productSizeRepo *productrepo.ProductSizeRepository) *ProductUseCase {
	return &ProductUseCase{
		DB:                       db,
		Log:                      log,
		Validate:                 validate,
		Cloudinary:               cloudinary,
		ProductRepository:        productRepo,
		CategoryRepository:       categoryRepository,
		ProductImageRepository:   productImageRepo,
		VariantImageRepository:   variantImageRepo,
		ProductVariantRepository: productVariantRepo,
		ProductSizeRepository:    productSizeRepo,
	}
}

func (c *ProductUseCase) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	category, err := c.CategoryRepository.ValidateCategoryExistence(tx, request.CategoryId)
	if err != nil {
		c.Log.Warnf("Category not found: %v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	subCategory, err := c.CategoryRepository.ValidateCategoryExistence(tx, request.SubCategoryId)
	if err != nil {
		c.Log.Warnf("Sub category not found: %v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "Sub category not found")
	}
	product := &entity.Product{
		StyleCode:     request.StyleCode,
		Name:          request.Name,
		Description:   request.Description,
		Gender:        request.Gender,
		CategoryId:    request.CategoryId,
		SubCategoryId: request.SubCategoryId,
		BasePrice:     request.BasePrice,
		IsVisible:     request.IsVisible,
		ReleaseDate:   request.ReleaseDate,
	}

	if err := c.ProductRepository.Create(tx, product); err != nil {
		c.Log.Warnf("Failed create product  to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	product.Category = category
	product.SubCategory = subCategory

	var productImages []*entity.ProductImage
	for _, image := range request.ProductImages {

		imageUrl, _ := pkg.UploadToCloudinary(c.Cloudinary, ctx, image.Image, product.ID, "")

		productImage := &entity.ProductImage{
			ProductId:    product.ID,
			ImageUrl:     imageUrl,
			DisplayOrder: image.DisplayOrder,
		}

		productImages = append(productImages, productImage)

	}

	if err := c.ProductImageRepository.BulkCreate(tx, productImages); err != nil {
		c.Log.Warnf("Failed create images  to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	var productVariants []*entity.ProductVariant
	var variantImages []*entity.VariantImage
	var productSizes []*entity.ProductSize
	for _, variant := range request.ProductVariants {
		productVariant := &entity.ProductVariant{
			ProductId:   product.ID,
			SKU:         variant.SKU,
			ColorName:   variant.ColorName,
			Weight:      variant.Weight,
			IsAvailable: variant.IsAvailable,
		}

		if err := c.ProductVariantRepository.Create(tx, productVariant); err != nil {
			c.Log.Warnf("Failed create variant  to database : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		for _, image := range variant.VariantImages {
			prefixVariantImg := fmt.Sprintf("%s/%s", productVariant.ColorName, productVariant.ID)
			imageUrl, _ := pkg.UploadToCloudinary(c.Cloudinary, ctx, image.Image, product.ID, prefixVariantImg)
			variantImage := &entity.VariantImage{
				VarianId:     productVariant.ID,
				ImageUrl:     imageUrl,
				DisplayOrder: image.DisplayOrder,
			}

			variantImages = append(variantImages, variantImage)
		}

		for _, size := range variant.ProductSizes {
			productSize := &entity.ProductSize{
				VariantId:     productVariant.ID,
				SKU:           size.SKU,
				Size:          size.Size,
				Discount:      size.Discount,
				Price:         size.Price,
				StockQuantity: size.StockQuantity,
				IsAvailable:   size.IsAvailable,
			}

			productSizes = append(productSizes, productSize)
		}

		if err := c.VariantImageRepository.BulkCreate(tx, variantImages); err != nil {
			c.Log.Warnf("Failed create images variant  to database : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		if err := c.ProductSizeRepository.BulkCreate(tx, productSizes); err != nil {
			c.Log.Warnf("Failed create size variant  to database : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		productVariants = append(productVariants, productVariant)
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return converter.ProductToResponse(product, productVariants, productImages, variantImages, productSizes), nil
}

func (c *ProductUseCase) Get(ctx context.Context, request *model.ProductGetRequest) (*model.ProductResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	product := new(entity.Product)
	if err := c.ProductRepository.FindDetailsProduct(tx, product, request.ID); err != nil {
		c.Log.Warnf("Product with id %s not found", request.ID)
		return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	// Parsing productVariant to pointer
	productVariants := make([]*entity.ProductVariant, len(product.ProductVariant))
	for i := range product.ProductVariant {
		productVariants[i] = &product.ProductVariant[i]
	}

	// Parsing productImages to pointer
	productImages := make([]*entity.ProductImage, len(product.ProductImage))
	for i := range product.ProductImage {
		productImages[i] = &product.ProductImage[i]
	}

	return converter.ProductToResponse(product, productVariants, productImages, getVariantImages(productVariants), getProductSizes(productVariants)), nil
}

func (c *ProductUseCase) List(ctx context.Context, request *model.ProductGetListRequest) ([]*model.ProductResponse, *model.PageMetadata, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)

		message := pkg.ParseValidationErrors(err)
		return nil, nil, fiber.NewError(fiber.StatusBadRequest, message)
	}

	products, pagination, err := c.ProductRepository.FindAllProducts(tx, request)
	fmt.Println("product: ", products)
	if err != nil {
		c.Log.Warnf("No products found for parameters: %+v", request)
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	// Parsing to pointer
	responses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		productVariants := make([]*entity.ProductVariant, len(product.ProductVariant))

		for j, variant := range product.ProductVariant {
			productVariants[j] = &variant
		}

		productImages := make([]*entity.ProductImage, len(product.ProductImage))
		for j, productImage := range product.ProductImage {
			productImages[j] = &productImage
		}

		responses[i] = converter.ProductToResponse(product, productVariants, productImages, getVariantImages(productVariants), getProductSizes(productVariants))
	}

	return responses, pagination, nil
}

func getVariantImages(productVariants []*entity.ProductVariant) []*entity.VariantImage {
	var variantImages []*entity.VariantImage
	for _, variant := range productVariants {
		for i := range variant.VariantImages {
			variantImages = append(variantImages, &variant.VariantImages[i])
		}
	}
	return variantImages
}

func getProductSizes(productVariants []*entity.ProductVariant) []*entity.ProductSize {
	var productSizes []*entity.ProductSize
	for _, variant := range productVariants {
		for i := range variant.ProductSizes {
			productSizes = append(productSizes, &variant.ProductSizes[i])
		}
	}
	return productSizes
}
