package tests_test

import (
	"testing"
	"trendora/config"

	"github.com/stretchr/testify/assert"
)

func TestConnectionDatabase(t *testing.T) {
	db := config.ConnectDB()
	assert.NotNil(t, db)

	sqlDB, err := db.DB()
	assert.Nil(t, err)
	defer sqlDB.Close()

	err = sqlDB.Ping()

	assert.Nil(t, err)
}
