package controller

import (
	"bytes"
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	//authentication
	tokenResponse := authenticationCreate()

	//service
	createProductRequest := model.ProductCreateOrUpdateModel{
		Name:     "Test Product",
		Price:    1000,
		Quantity: 1000,
	}
	requestBody, _ := json.Marshal(createProductRequest)

	request := httptest.NewRequest("POST", "/v1/api/product", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+tokenResponse["token"].(string))

	response, _ := appTest.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)

	webResponse := model.GeneralResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "Success", webResponse.Message)

	jsonData, _ := json.Marshal(webResponse.Data)
	createProductResponse := model.ProductCreateOrUpdateModel{}
	_ = json.Unmarshal(jsonData, &createProductResponse)

	assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
	assert.Equal(t, createProductRequest.Price, createProductResponse.Price)
	assert.Equal(t, createProductRequest.Quantity, createProductResponse.Quantity)
}
