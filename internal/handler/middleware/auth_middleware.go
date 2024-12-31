package middleware

import (
	"strings"

	"github.com/bwafi/trendora-backend/internal/model"
	adminusecase "github.com/bwafi/trendora-backend/internal/usecase/admin"
	customerusecase "github.com/bwafi/trendora-backend/internal/usecase/customer"
	"github.com/bwafi/trendora-backend/pkg"
	"github.com/gofiber/fiber/v3"
)

func CustomerAuthMiddleware(customerCase *customerusecase.CustomerUseCase) fiber.Handler {
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

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			customerCase.Log.Warn("Authorization header must be in format: Bearer <token>")
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Authorization header format must be 'Bearer <token>'",
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

		if jwtClaims.Role != "customer" {
			customerCase.Log.Warnf("Access denied for role: %s", jwtClaims.Role)

			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Access denied. This route is restricted to customers only.",
				},
			})
		}

		auth := &model.Auth{
			ID:   jwtClaims.Subject,
			Role: jwtClaims.Role,
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func AdminAuthMiddleware(adminCase *adminusecase.AdminUseCase) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		stringToken := ctx.Get(fiber.HeaderAuthorization, "")

		if stringToken == "" {
			adminCase.Log.Warn("Missing authentication token")
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Missing authentication token",
				},
			})
		}

		adminCase.Log.Debugf("Authorization header received: %s...", stringToken[:len(stringToken)/2])

		tokenParts := strings.Split(stringToken, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			adminCase.Log.Warn("Authorization header must be in format: Bearer <token>")
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Authorization header format must be 'Bearer <token>'",
				},
			})
		}

		token := tokenParts[1]
		secretKey := adminCase.Config.GetString("jwt.accessToken")

		jwtClaims, err := pkg.VerifyToken(token, adminCase.Log, secretKey)
		if err != nil {
			adminCase.Log.Warnf("Failed to verify token: %v", err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Invalid or expired token",
				},
			})
		}

		if jwtClaims.Role != "admin" {
			adminCase.Log.Warnf("Access denied for role: %s", jwtClaims.Role)

			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.CustomerResponse]{
				Errors: &model.ErrorResponse{
					Code:    fiber.StatusUnauthorized,
					Message: "Access denied. This route is restricted to admins only.",
				},
			})
		}

		auth := &model.Auth{
			ID:   jwtClaims.Subject,
			Role: jwtClaims.Role,
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
