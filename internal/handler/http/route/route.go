package route

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App                *fiber.App
	CustomerController *http.CustomerController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/customers/register", c.CustomerController.Register)
	c.App.Post("/api/customers/login", c.CustomerController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Patch("/api/auth/customers", c.CustomerController.Update)
	c.App.Delete("/api/auth/customers", c.CustomerController.Delete)
}
