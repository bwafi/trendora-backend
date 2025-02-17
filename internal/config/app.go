package config

import (
	"github.com/bwafi/trendora-backend/internal/handler/http"
	"github.com/bwafi/trendora-backend/internal/handler/http/route"
	"github.com/bwafi/trendora-backend/internal/handler/middleware"
	adminrepo "github.com/bwafi/trendora-backend/internal/repository/admin"
	cartrepo "github.com/bwafi/trendora-backend/internal/repository/cart"
	customerrepo "github.com/bwafi/trendora-backend/internal/repository/customer"
	productrepo "github.com/bwafi/trendora-backend/internal/repository/product"
	adminusecase "github.com/bwafi/trendora-backend/internal/usecase/admin"
	cartusecase "github.com/bwafi/trendora-backend/internal/usecase/cart"
	customerusecase "github.com/bwafi/trendora-backend/internal/usecase/customer"
	productusecase "github.com/bwafi/trendora-backend/internal/usecase/product"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	App        *fiber.App
	Log        *logrus.Logger
	Validate   *validator.Validate
	Config     *viper.Viper
	Cloudinary *cloudinary.Cloudinary
}

func Bootstrap(config *BootstrapConfig) {
	// Repository
	customerRepository := customerrepo.NewCustomerRepository(config.Log)
	customerSessionRepository := customerrepo.NewCustomerSessionRepository(config.Log)
	customerAddressRepository := customerrepo.NewCustomerAddressesRepository(config.Log)

	productRepository := productrepo.NewProductRepository(config.Log)
	categoryRepository := productrepo.NewCategoryRepository(config.Log)
	productImageRepository := productrepo.NewProductImageRepository(config.Log)
	variantImageRepository := productrepo.NewVariantImageRepository(config.Log)
	productVariantRepository := productrepo.NewProductVariantRepository(config.Log)
	productSizeRepository := productrepo.NewProductSizeRepository(config.Log)
	productReviewRepository := productrepo.NewProductReviewRepository(config.Log)

	cartItemRepo := cartrepo.NewCartItemRepository(config.Log)

	adminRepository := adminrepo.NewAdminrRepository(config.Log)

	// Usecase
	customerUseCase := customerusecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, config.Config, customerRepository, customerSessionRepository)
	customerAddressUseCase := customerusecase.NewCustomerAddressUsecase(config.DB, config.Log, config.Validate, config.Config, customerAddressRepository)
	productUseCase := productusecase.NewProductUseCase(config.DB, config.Log, config.Validate, config.Cloudinary, productRepository, categoryRepository, productImageRepository, variantImageRepository, productVariantRepository, productSizeRepository)
	cartItemUseCase := cartusecase.NewCartItemUseCase(config.DB, config.Log, config.Validate, cartItemRepo, customerRepository, productRepository, productVariantRepository)
	adminUseCase := adminusecase.NewAdminUseCase(config.DB, config.Log, config.Validate, config.Config, adminRepository)
	productReviewUseCase := productusecase.NewProductReviewUsecase(config.DB, config.Log, config.Validate, productReviewRepository, productRepository, customerRepository)

	customerController := http.NewCustomerController(customerUseCase, config.Log)
	cusomerAddressController := http.NewCustomerAddressController(customerAddressUseCase, config.Log, config.Config)
	productController := http.NewProductController(productUseCase, config.Log)
	cartItemController := http.NewCartItemController(config.Log, cartItemUseCase)
	adminController := http.NewAdminUseCase(adminUseCase, config.Log)
	productReviewController := http.NewProductReviewController(productReviewUseCase, config.Log)

	customerAuthMiddleware := middleware.CustomerAuthMiddleware(customerUseCase)
	adminAuthMiddleware := middleware.AdminAuthMiddleware(adminUseCase)

	routeConfig := route.RouteConfig{
		App:                       config.App,
		CustomerController:        customerController,
		CustomerAddressController: cusomerAddressController,
		ProductController:         productController,
		AdminController:           adminController,
		CartItemController:        cartItemController,
		ProductReviewController:   productReviewController,
		CustomerAuthMiddleware:    customerAuthMiddleware,
		AdminAuthMiddleware:       adminAuthMiddleware,
	}

	routeConfig.Setup()
}
