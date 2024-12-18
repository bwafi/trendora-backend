package repository

import (
	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/sirupsen/logrus"
)

type CustomerSessionRepository struct {
	Repository[entity.CustomerSessions]
	Log *logrus.Logger
}

func NewCustomerSessionRepository(log *logrus.Logger) *CustomerSessionRepository {
	return &CustomerSessionRepository{
		Log: log,
	}
}
