package converter

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
)

func AdminToAdminResponse(admin *entity.Admin) *model.AdminResponse {
	return &model.AdminResponse{
		ID:        admin.ID,
		Name:      admin.Name,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func AdminToAuthResponse(admin *entity.Admin, accessToken string, refreshToken string) *model.AdminResponse {
	return &model.AdminResponse{
		ID:           admin.ID,
		Name:         admin.Name,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		CreatedAt:    admin.CreatedAt,
		UpdatedAt:    admin.UpdatedAt,
	}
}
