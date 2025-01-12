package route

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App                       *fiber.App
	CustomerController        *http.CustomerController
	CustomerAddressController *http.CustomerAddressController
	ProductController         *http.ProductController
	AdminController           *http.AdminController
	CartItemController        *http.CartItemController
	CustomerAuthMiddleware    fiber.Handler
	AdminAuthMiddleware       fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
	c.SetupAdminRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	customerRoutes := c.App.Group("/api/customers")
	customerRoutes.Post("/register", c.CustomerController.Register)
	customerRoutes.Post("/login", c.CustomerController.Login)

	adminRoutes := c.App.Group("/api/admins")
	adminRoutes.Post("/register", c.AdminController.Register)
	adminRoutes.Post("/login", c.AdminController.Login)

	// Public Product Routes (Accessible to both Admin and Customer)
	// productRoutes := c.App.Group("/api/products")
	// productRoutes.Get("/", c.ProductController.List)          // List Products (Shared Route)
	// productRoutes.Get("/:productId", c.ProductController.Get) // View a specific Product (Shared Route)
}

func (c *RouteConfig) SetupAuthRoute() {
	customerAuthRoutes := c.App.Group("/api/customers", c.CustomerAuthMiddleware)
	customerAuthRoutes.Patch("/", c.CustomerController.Update)
	customerAuthRoutes.Delete("/", c.CustomerController.Delete)

	// Address route
	addressRoutes := customerAuthRoutes.Group("/address")
	addressRoutes.Post("/", c.CustomerAddressController.Create)
	addressRoutes.Get("/", c.CustomerAddressController.List)
	addressRoutes.Get("/:addressId", c.CustomerAddressController.Get)
	addressRoutes.Patch("/:addressId", c.CustomerAddressController.Update)
	addressRoutes.Delete("/:addressId", c.CustomerAddressController.Delete)

	// Cart route
	cartItemRoute := customerAuthRoutes.Group("/carts")
	cartItemRoute.Post("/", c.CartItemController.Create)
	cartItemRoute.Get("/:cartId", c.CartItemController.Get)
	cartItemRoute.Patch("/:cartId", c.CartItemController.Update)
	cartItemRoute.Delete("/:cartId", c.CartItemController.Delete)
}

func (c *RouteConfig) SetupAdminRoute() {
	// Admin-only Routes
	adminRoutes := c.App.Group("/api/admins", c.AdminAuthMiddleware)

	// Admin product routes (for creating and managing products)
	adminProductRoutes := adminRoutes.Group("/products")
	adminProductRoutes.Post("/", c.ProductController.Create) // Create Product (Admin-only)
}
