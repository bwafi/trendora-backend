package productrepo

import (
	"errors"
	"fmt"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	repository.Repository[entity.Category]
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Log: log,
	}
}

func (c *CategoryRepository) ValidateCategoryExistence(tx *gorm.DB, id string) (*entity.Category, error) {
	category := new(entity.Category)
	if err := tx.First(category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("category with ID %s not found", id)
		}
		return nil, err
	}

	return category, nil
}
