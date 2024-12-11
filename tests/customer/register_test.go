package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/bwafi/trendora-backend/tests"
	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccessWithEmail(t *testing.T) {
	tests.ClearAll()

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
		Name:         "Syahroni",
		Password:     "rahasia",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, "Customer registration successful", responseBody.Message)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestRegisterSuccessWithPhone(t *testing.T) {
	tests.ClearAll()

	requestBody := &model.CustomerRegisterRequest{
		PhoneNumber: tests.StrPointer("082332323232"),
		Name:        "Syahroni",
		Password:    "rahasia",
		Gender:      true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, "Customer registration successful", responseBody.Message)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestRegisterFailsWhenDuplicatePhone(t *testing.T) {
	tests.ClearAll()
	TestRegisterSuccessWithPhone(t)

	requestBody := &model.CustomerRegisterRequest{
		PhoneNumber: tests.StrPointer("082332323232"),
		Name:        "Syahroni",
		Password:    "rahasia",
		Gender:      true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
	assert.Equal(t, "Phone number already in use", responseBody.Errors.Message)
}

func TestRegisterFailsWhenDuplicateEmail(t *testing.T) {
	tests.ClearAll()
	TestRegisterSuccessWithEmail(t)

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
		Name:         "Syahroni",
		Password:     "rahasia",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
	assert.Equal(t, "Email already in use", responseBody.Errors.Message)
}

func TestRegisterFailsDuplicateEmail(t *testing.T) {
	tests.ClearAll()
	TestRegisterSuccessWithEmail(t)

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
		Name:         "Syahroni",
		Password:     "rahasia",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
	assert.Equal(t, "Email already in use", responseBody.Errors.Message)
}

func TestRegisterFailsWithoutPhoneOrEmail(t *testing.T) {
	tests.ClearAll()

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: nil,
		PhoneNumber:  nil,
		Name:         "Syahroni",
		Password:     "rahasia",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Phone number or Email is required", responseBody.Errors.Message)
}

func TestRegisterFailsWithoutName(t *testing.T) {
	tests.ClearAll()

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
		Name:         "",
		Password:     "rahasia",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Name is required", responseBody.Errors.Message)
}

func TestRegisterFailsWithoutPassword(t *testing.T) {
	tests.ClearAll()

	requestBody := &model.CustomerRegisterRequest{
		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
		Name:         "Syahroni",
		Password:     "",
		Gender:       true,
	}

	requestJson, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := tests.App.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[*model.CustomerResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Password is required", responseBody.Errors.Message)
}

// TODO: Test Fails validation min or max character

// func TestRegisterFailsMinName(t *testing.T) {
// 	tests.ClearAll()
//
// 	requestBody := &model.CustomerRegisterRequest{
// 		EmailAddress: tests.StrPointer("syahroni@gmail.com"),
// 		Name:         "Sy",
// 		Password:     "rahasi",
// 		Gender:       true,
// 	}
//
// 	requestJson, _ := json.Marshal(requestBody)
//
// 	request := httptest.NewRequest(http.MethodPost, "/api/customers/register", strings.NewReader(string(requestJson)))
// 	request.Header.Set("Content-Type", "application/json")
// 	request.Header.Set("Accept", "application/json")
//
// 	response, err := tests.App.Test(request)
// 	assert.Nil(t, err)
//
// 	responseBytes, err := io.ReadAll(response.Body)
// 	assert.Nil(t, err)
//
// 	responseBody := new(model.WebResponse[*model.CustomerResponse])
// 	err = json.Unmarshal(responseBytes, responseBody)
// 	assert.Nil(t, err)
//
// 	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
// 	assert.Equal(t, "Password is required", responseBody.Errors.Message)
// }
