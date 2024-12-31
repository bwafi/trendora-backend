package adminrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AdminRepository struct {
	repository.Repository[entity.Admin]
	Log *logrus.Logger
}

func NewAdminrRepository(log *logrus.Logger) *AdminRepository {
	return &AdminRepository{
		Log: log,
	}
}

func (c *AdminRepository) ExistsByEmail(tx *gorm.DB, email *string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Admin{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *AdminRepository) ExistsByPhoneNumber(tx *gorm.DB, phoneNumber *string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Admin{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
