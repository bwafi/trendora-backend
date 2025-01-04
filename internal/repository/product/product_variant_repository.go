package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductVariantRepository struct {
	repository.Repository[entity.ProductVariant]
	Log *logrus.Logger
}

func NewProductVariantRepository(log *logrus.Logger) *ProductVariantRepository {
	return &ProductVariantRepository{
		Log: log,
	}
}
