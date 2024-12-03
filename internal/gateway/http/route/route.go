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
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.CustomerController.Register)
}
