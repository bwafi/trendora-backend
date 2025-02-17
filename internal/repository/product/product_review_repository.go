package productrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type ProductReviewRepository struct {
	repository.Repository[entity.ProductReview]
	Log *logrus.Logger
}

func NewProductReviewRepository(log *logrus.Logger) *ProductReviewRepository {
	return &ProductReviewRepository{
		Log: log,
	}
}
