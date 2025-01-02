package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductImageRepository struct {
	repository.Repository[entity.ProductImage]
	Log *logrus.Logger
}

func NewProductImageRepository(log *logrus.Logger) *ProductImageRepository {
	return &ProductImageRepository{
		Log: log,
	}
}
