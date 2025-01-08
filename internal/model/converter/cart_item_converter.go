package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func CartItemToReponse(cartItem *entity.CartItem) *model.CartItemResponse {
	return &model.CartItemResponse{
		ID:         cartItem.ID,
		CustomerId: cartItem.CustomerId,
		ProductId:  cartItem.ProductId,
		VariantId:  cartItem.VariantId,
		Quantity:   cartItem.Quantity,
		CreatedAt:  cartItem.CreatedAt,
		UpdatedAt:  cartItem.UpdatedAt,
	}
}
