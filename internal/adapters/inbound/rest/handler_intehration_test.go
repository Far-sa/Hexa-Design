package handler_test

import (
	handler "hexa-design/internal/adapters/inbound/rest"
	repository "hexa-design/internal/adapters/outbound/repositories"
	"hexa-design/internal/domain/model"
	"hexa-design/internal/domain/services"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts_integration(t *testing.T) {
	t.Run("", func(t *testing.T) {
		
		//* Arrange
		expected := ""
		productRepo := repository.NewProductRepositoryMock()
		productRepo.On("GetProducts").Return(model.Product{
			ID:       271,
			Name:     "Product271",
			Quantity: 99,
		}, nil)

		productSvc := services.NewProductService(productRepo)
		productHandler := handler.NewProductHandler(productSvc)

		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		req := httptest.NewRequest("GET", "/products", nil)

		//* Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, expected, body)
		}
	})
}
