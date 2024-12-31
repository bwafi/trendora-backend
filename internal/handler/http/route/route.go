package route

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App                       *fiber.App
	CustomerController        *http.CustomerController
	CustomerAddressController *http.CustomerAddressController
	AdminController           *http.AdminController
	CustomerAuthMiddleware    fiber.Handler
	AdminAuthMiddleware       fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	customerRoutes := c.App.Group("/api/customers")
	customerRoutes.Post("/register", c.CustomerController.Register)
	customerRoutes.Post("/login", c.CustomerController.Login)

	adminRoutes := c.App.Group("/api/admins")
	adminRoutes.Post("/register", c.AdminController.Register)
	adminRoutes.Post("/login", c.AdminController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	customerAuthRoutes := c.App.Group("/api/customers", c.CustomerAuthMiddleware)
	customerAuthRoutes.Patch("/", c.CustomerController.Update)
	customerAuthRoutes.Delete("/", c.CustomerController.Delete)

	addressRoutes := customerAuthRoutes.Group("/address")
	addressRoutes.Post("/", c.CustomerAddressController.Create)
	addressRoutes.Get("/", c.CustomerAddressController.List)
	addressRoutes.Get("/:addressId", c.CustomerAddressController.Get)
	addressRoutes.Patch("/:addressId", c.CustomerAddressController.Update)
	addressRoutes.Delete("/:addressId", c.CustomerAddressController.Delete)
}
