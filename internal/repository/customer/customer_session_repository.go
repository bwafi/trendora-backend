package customerrepo

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type CustomerSessionRepository struct {
	repository.Repository[entity.CustomerSessions]
	Log *logrus.Logger
}

func NewCustomerSessionRepository(log *logrus.Logger) *CustomerSessionRepository {
	return &CustomerSessionRepository{
		Log: log,
	}
}
