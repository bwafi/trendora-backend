package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type VariantImageRepository struct {
	repository.Repository[entity.VariantImage]
	Log *logrus.Logger
}

func NewVariantImageRepository(log *logrus.Logger) *VariantImageRepository {
	return &VariantImageRepository{
		Log: log,
	}
}
