package config

import (
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/go-playground/validator/v10"
)

func NewValidation() *validator.Validate {
	return validator.New()
}

func MustValidRegisterCustomer(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(model.CustomerRegisterRequest)

	if registerRequest.PhoneNumber != "" || registerRequest.EmailAddress != "" {
		return
	}

	level.ReportError(registerRequest.PhoneNumber, "PhoneNumber", "phone_number", "required_without_email", "")
	level.ReportError(registerRequest.EmailAddress, "EmailAddress", "email_address", "required_without_phone", "")
}
