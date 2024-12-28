package customerrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	repository.Repository[entity.Customers]
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}

func (c *CustomerRepository) FindById(tx *gorm.DB, customer *entity.Customers, id *string) error {
	return tx.Where("id = ?", id).Take(customer).Error
}

func (c *CustomerRepository) FindByEmailOrPhone(tx *gorm.DB, customer *entity.Customers, email *string, phone *string) error {
	return tx.Where("email_address = ?", email).Or("phone_number = ?", phone).Take(customer).Error
}

func (c *CustomerRepository) ExistsByEmail(tx *gorm.DB, email *string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Customers{}).Where("email_address = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *CustomerRepository) ExistsByPhoneNumber(tx *gorm.DB, phoneNumber *string) (int64, error) {
	var count int64
	err := tx.Model(&entity.Customers{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
