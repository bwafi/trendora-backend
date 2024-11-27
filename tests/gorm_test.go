package tests

import (
	"testing"

	"github.com/bwafi/trendora-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectionToDB(t *testing.T) {
	viper := config.NewViper()
	logrus := config.NewLogger(viper)
	db := config.NewDatabase(viper, logrus)

	assert.NotNil(t, db)

	connection, err := db.DB()
	assert.Nil(t, err)

	err = connection.Ping()

	assert.Nil(t, err)
}
