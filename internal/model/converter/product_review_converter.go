package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func ProductReviewToResponse(productReview *entity.ProductReview) *model.ProductReviewResponse {
	return &model.ProductReviewResponse{
		ID:         productReview.ID,
		ProductId:  productReview.ProductId,
		CustomerID: productReview.CustomerID,
		Rating:     productReview.Rating,
		Comment:    productReview.Comment,
		CreatedAt:  productReview.CreatedAt,
		UpdatedAt:  productReview.UpdatedAt,
	}
}
