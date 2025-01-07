package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductSizeRepository struct {
	repository.Repository[entity.ProductSize]
	Log *logrus.Logger
}

func NewProductSizeRepository(log *logrus.Logger) *ProductSizeRepository {
	return &ProductSizeRepository{
		Log: log,
	}
}
