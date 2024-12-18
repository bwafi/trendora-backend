package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func CustomerToResponse(customer *entity.Customers) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

func CustomerToAuthResponse(customer *entity.Customers, accessToken string, refreshToken string) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:           customer.ID,
		Name:         customer.Name,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		CreatedAt:    customer.CreatedAt,
		UpdatedAt:    customer.UpdatedAt,
	}
}

func CustomerAddressToResponse(customer *entity.CustomerAddresses) *model.AddressResponse {
	return &model.AddressResponse{
		ID:            customer.ID,
		RecipientName: customer.CustomerID,
		PhoneNumber:   customer.PhoneNumber,
		AddressType:   customer.AddressType,
		City:          customer.City,
		Province:      customer.Province,
		SubDistrict:   customer.SubDistrict,
		PostalCode:    customer.PostalCode,
		CourierNotes:  customer.CourierNotes,
		CreatedAt:     customer.CreatedAt,
		UpdatedAt:     customer.UpdatedAt,
	}
}
