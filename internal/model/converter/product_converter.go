package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func ProductToResponse(product *entity.Product, productVariants []*entity.ProductVariant, productImages []*entity.ProductImage, variantImages []*entity.VariantImage, productSizes []*entity.ProductSize) *model.ProductResponse {
	return &model.ProductResponse{
		ID:            product.ID,
		StyleCode:     product.StyleCode,
		Name:          product.Name,
		Description:   product.Description,
		Gender:        product.Gender,
		CategoryId:    product.CategoryId,
		SubCategoryId: product.SubCategoryId,
		BasePrice:     product.BasePrice,
		IsVisible:     product.IsVisible,
		ReleaseDate:   product.ReleaseDate,
		Category: model.CategoryResponse{
			ID:       product.SubCategory.ID,
			ParentId: product.SubCategory.ParentId,
			Name:     product.SubCategory.Name,
			Slug:     product.SubCategory.Slug,
			ParentCategory: &model.CategoryResponse{
				ID:   product.Category.ID,
				Name: product.Category.Name,
				Slug: product.Category.Slug,
			},
		},
		ProductVariant: ConvertToProductVariantResponse(productVariants, variantImages, productSizes),
		ProductImages:  ConvertToProductImageResponse(productImages),
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}
}

func ConvertToProductImageResponse(productImages []*entity.ProductImage) []*model.ImageResponse {
	productImageSlice := make([]*model.ImageResponse, len(productImages))
	for i, image := range productImages {
		productImageSlice[i] = &model.ImageResponse{
			ID:           image.ID,
			ProductId:    image.ProductId,
			ImageUrl:     image.ImageUrl,
			DisplayOrder: image.DisplayOrder,
		}
	}

	return productImageSlice
}

func ConvertToProductVariantResponse(productVariants []*entity.ProductVariant, variantImages []*entity.VariantImage, productSizes []*entity.ProductSize) []*model.ProductVariantResponse {
	productVariantSlice := make([]*model.ProductVariantResponse, len(productVariants))
	for i, variant := range productVariants {
		productVariantSlice[i] = &model.ProductVariantResponse{
			ID:            variant.ID,
			ProductId:     variant.ProductId,
			SKU:           variant.SKU,
			ColorName:     variant.ColorName,
			Weight:        variant.Weight,
			IsAvailable:   variant.IsAvailable,
			VariantImages: ConvertToVariantImageResponse(variantImages),
			ProductSizes:  ConvertToProductSizeResponse(productSizes),
		}
	}

	return productVariantSlice
}

func ConvertToVariantImageResponse(variantImages []*entity.VariantImage) []*model.ImageResponse {
	variantImageSlice := make([]*model.ImageResponse, len(variantImages))
	for i, image := range variantImages {
		variantImageSlice[i] = &model.ImageResponse{
			ID:           image.ID,
			VarianId:     image.VarianId,
			ImageUrl:     image.ImageUrl,
			DisplayOrder: image.DisplayOrder,
		}
	}

	return variantImageSlice
}

func ConvertToProductSizeResponse(productSizes []*entity.ProductSize) []*model.ProductSizeResponse {
	productSizeSlice := make([]*model.ProductSizeResponse, len(productSizes))

	for i, size := range productSizes {
		productSizeSlice[i] = &model.ProductSizeResponse{
			ID:            size.ID,
			VariantId:     size.VariantId,
			SKU:           size.SKU,
			Size:          size.Size,
			Discount:      size.Discount,
			Price:         size.Price,
			StockQuantity: size.StockQuantity,
			IsAvailable:   size.IsAvailable,
		}
	}

	return productSizeSlice
}
