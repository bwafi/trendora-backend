package config

import (
	"github.com/bwafi/trendora-backend/internal/model"
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
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		} else {
			message = err.Error()
		}

		return ctx.Status(code).JSON(model.WebResponse[*string]{
			Errors: &model.ErrorResponse{
				Code:    code,
				Message: message,
			},
		})
	}
}
