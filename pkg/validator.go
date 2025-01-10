package pkg

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func MustValidRegisterCustomer(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(model.CustomerRegisterRequest)

	if registerRequest.PhoneNumber != nil || registerRequest.EmailAddress != nil {
		return
	}

	level.ReportError(registerRequest.PhoneNumber, "PhoneNumber", "phone_number", "required_without_email", "")
	level.ReportError(registerRequest.EmailAddress, "EmailAddress", "email_address", "required_without_phone", "")
}

func ParseValidationErrors(err error) string {
	var message string

	if validatorErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validatorErrors {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", CamelCaseToReadable(err.Field()))
			case "email":
				message = fmt.Sprintf("%s is not valid email", CamelCaseToReadable(err.Field()))
			case "required_without_phone", "required_without_email":
				message = "Phone number or Email is required"
			case "max":
				message = fmt.Sprintf("%s must be less than or equal to %s", CamelCaseToReadable(err.Field()), err.Param())
			case "min":
				message = fmt.Sprintf("%s must be greater than or equal to %s", CamelCaseToReadable(err.Field()), err.Param())
			case "gte":
				message = fmt.Sprintf("%s must be greater than 0 or positive number", CamelCaseToReadable(err.Field()))
			default:
				message = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}
	}

	return message
}

func CamelCaseToReadable(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")

	output := re.ReplaceAllString(input, "${1} ${2}")

	words := strings.Split(cases.Lower(language.Tag{}).String(output), " ")

	if len(words) > 0 {
		words[0] = cases.Title(language.Tag{}).String(words[0])
	}

	return strings.Join(words, " ")
}
