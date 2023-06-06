package handler_test

import (
	handler "hexa-design/internal/adapters/inbound/rest"
	"hexa-design/internal/domain/services"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		//* Arrange
		expected := ""
		productsSvc := services.NewProductServiceMock()
		productsSvc.On("GetProducts").Return(expected, nil)

		productHandler := handler.NewProductHandler(productsSvc)

		//* http://localhost:8000/products
		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		req := httptest.NewRequest("GET", "/products", nil)

		//* Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//* Assert
		assert.Equal(t, fiber.StatusOK, res.StatusCode)
	})
}
