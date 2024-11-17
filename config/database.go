package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         string
}

func LoadDatabase() *Database {
	return &Database{
		Host:         GetEnv("HOST_DB", "localhost"),
		User:         GetEnv("USER_DB", "postgres"),
		Password:     GetEnv("PASSWORD_DB", ""),
		DatabaseName: GetEnv("DATABASE_NAME", ""),
		Port:         GetEnv("PORT_DB", "5432"),
	}
}

func ConnectDB() *gorm.DB {
	dbConfig := LoadDatabase()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DatabaseName, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get *sql.DB: %v", err))
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	return db
}
