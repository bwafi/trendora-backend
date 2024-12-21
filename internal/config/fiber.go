package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      viper.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		// Default status code dan message
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		// Jika error merupakan instance dari fiber.Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		} else {
			// Cek error custom lainnya (opsional)
			message = err.Error()
		}

		// Kembalikan response JSON dengan struktur yang diminta
		return ctx.Status(code).JSON(fiber.Map{
			"errors": fiber.Map{
				"status":  code,
				"message": message,
			},
		})
	}
}
