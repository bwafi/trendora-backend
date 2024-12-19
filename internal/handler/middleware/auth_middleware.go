package middleware

import (
	"strings"

	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/internal/usecase"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(customerCase *usecase.CustomerUseCase) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		stringToken := ctx.Get(fiber.HeaderAuthorization, "")

		if stringToken == "" {
			customerCase.Log.Warn("Missing authentication token")
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Missing authentication token",
				},
			})
		}

		customerCase.Log.Debugf("Authorization header received: %s...", stringToken[:len(stringToken)/2])

		tokenParts := strings.Split(stringToken, " ")

		if len(tokenParts) != 2 && tokenParts[0] != "Bearer" {
			customerCase.Log.Warn("Invalid authentication token")
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Invalid authentication token",
				},
			})
		}

		token := tokenParts[1]
		secretKey := customerCase.Config.GetString("jwt.accessToken")

		jwtClaims, err := pkg.VerifyToken(token, customerCase.Log, secretKey)
		if err != nil {
			customerCase.Log.Warnf("Failed to verify token: %v", err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Invalid or expired token",
				},
			})
		}

		ctx.Locals("auth", jwtClaims.ID)
		return ctx.Next()
	}
}

func GetUser(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
