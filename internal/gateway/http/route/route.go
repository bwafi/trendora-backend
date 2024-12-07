package route

import (
	"github.com/bwafi/trendora-backend/internal/gateway/http"
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
	c.App.Post("/api/customers", c.CustomerController.Register)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Patch("/api/customers", c.CustomerController.Update)
}
