package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func CartItemToReponse(cartItem *entity.CartItem) *model.CartItemResponse {
	// productVariants := make([]model.ProductVariantResponse, len(cartItem.Product.ProductVariant))
	//
	// for i, variant := range cartItem.Product.ProductVariant {
	// 	variantImages := make([]model.ImageResponse, len(variant.VariantImages))
	//
	// 	for i, image := range variant.VariantImages {
	// 		variantImages[i] = model.ImageResponse{
	// 			ID:           image.ID,
	// 			VarianId:     image.VarianId,
	// 			ImageUrl:     image.ImageUrl,
	// 			DisplayOrder: image.DisplayOrder,
	// 		}
	// 	}
	//
	// 	productVariants[i] = model.ProductVariantResponse{
	// 		ID:            variant.ID,
	// 		ProductId:     variant.ProductId,
	// 		SKU:           variant.SKU,
	// 		ColorName:     variant.ColorName,
	// 		Weight:        variant.Weight,
	// 		IsAvailable:   variant.IsAvailable,
	// 		VariantImages: []model.ImageResponse{},
	// 		ProductSizes:  []model.ProductSizeResponse{},
	// 	}
	// }

	return &model.CartItemResponse{
		ID:         cartItem.ID,
		CustomerId: cartItem.CustomerId,
		ProductId:  cartItem.ProductId,
		VariantId:  cartItem.VariantId,
		Quantity:   cartItem.Quantity,
		CreatedAt:  cartItem.CreatedAt,
		UpdatedAt:  cartItem.UpdatedAt,
		Product: &model.ProductResponse{
			ID: cartItem.Product.ID,
			// StyleCode:     cartItem.Product.StyleCode,
			// Name:          cartItem.Product.Name,
			// Description:   cartItem.Product.Description,
			// Gender:        cartItem.Product.Gender,
			// CategoryId:    cartItem.Product.CategoryId,
			// SubCategoryId: cartItem.Product.SubCategoryId,
			// BasePrice:     cartItem.Product.BasePrice,
			// IsVisible:     cartItem.Product.IsVisible,
			// ReleaseDate:   cartItem.Product.ReleaseDate,
			// Category: model.CategoryResponse{
			// 	ID:       cartItem.Product.SubCategory.ID,
			// 	ParentId: cartItem.Product.SubCategory.ParentId,
			// 	Name:     cartItem.Product.SubCategory.Name,
			// 	Slug:     cartItem.Product.SubCategory.Slug,
			// 	ParentCategory: &model.CategoryResponse{
			// 		ID:   cartItem.Product.Category.ID,
			// 		Name: cartItem.Product.Category.Name,
			// 		Slug: cartItem.Product.Category.Slug,
			// 	},
			// },
			// // ProductVariant: productVariants,
			// ProductImages: []model.ImageResponse{},
			// CreatedAt:     cartItem.Product.CreatedAt,
			// UpdatedAt:     cartItem.Product.UpdatedAt,
		},
	}
}
