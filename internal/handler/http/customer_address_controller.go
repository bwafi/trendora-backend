package http

import (
	"fmt"

	"github.com/bwafi/trendora-backend/internal/handler/middleware"
	"github.com/bwafi/trendora-backend/internal/model"
	customerusecase "github.com/bwafi/trendora-backend/internal/usecase/customer"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CustomerAddressController struct {
	Log                 *logrus.Logger
	CustomerAddressCase *customerusecase.CustomerAddressUsecase
	Config              *viper.Viper
}

func NewCustomerAddressController(useCase *customerusecase.CustomerAddressUsecase, logger *logrus.Logger, viper *viper.Viper) *CustomerAddressController {
	return &CustomerAddressController{
		Log:                 logger,
		CustomerAddressCase: useCase,
		Config:              viper,
	}
}

func (c *CustomerAddressController) Create(ctx fiber.Ctx) error {
	request := new(model.CreateAddressRequest)

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.AddressResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerAddressCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.AddressResponse]{
		Data: response,
	})
}

func (c *CustomerAddressController) Get(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	addressId := ctx.Params("addressId")
	fmt.Println("customer id nih", auth)

	request := &model.GetAddressRequest{
		ID:         addressId,
		CustomerID: auth.ID,
	}

	response, err := c.CustomerAddressCase.Get(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.AddressResponse]{
		Data: response,
	})
}

func (c *CustomerAddressController) List(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetAddressRequest{
		CustomerID: auth.ID,
	}

	response, err := c.CustomerAddressCase.List(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]model.AddressResponse]{
		Data: response,
	})
}

func (c *CustomerAddressController) Update(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	addressId := ctx.Params("addressId")
	request := new(model.UpdateAddressRequest)

	request.CustomerID = auth.ID
	request.ID = addressId

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.AddressResponse]{
			Errors: &model.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request body",
			},
		})
	}

	response, err := c.CustomerAddressCase.Update(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.AddressResponse]{
		Data: response,
	})
}

func (c *CustomerAddressController) Delete(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	addressId := ctx.Params("addressId")

	request := &model.DeleteAddressRequest{
		ID:         addressId,
		CustomerID: auth.ID,
	}

	err := c.CustomerAddressCase.Delete(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Address successfully deleted",
	})
}
