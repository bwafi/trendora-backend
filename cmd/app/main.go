package main

import (
	"fmt"
	"log"

	"github.com/bwafi/trendora-backend/internal/config"
	"github.com/gofiber/fiber/v3"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	logger := config.NewLogger(viper)
	db := config.NewDatabase(viper, logger)
	validate := config.NewValidation()
	cloudinary := config.NewCloudinary(viper)

	config.Bootstrap(&config.BootstrapConfig{
		DB:         db,
		App:        app,
		Log:        logger,
		Validate:   validate,
		Config:     viper,
		Cloudinary: cloudinary,
	})

	webPort := viper.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort), fiber.ListenConfig{EnablePrefork: viper.GetBool("web.prefork")})
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}
