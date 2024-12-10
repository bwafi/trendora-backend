package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func CustomerToResponse(customer *entity.Customers) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Token:     customer.Token,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}
