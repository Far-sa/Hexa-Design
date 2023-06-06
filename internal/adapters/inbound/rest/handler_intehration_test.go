package handler_test

import (
	handler "hexa-design/internal/adapters/inbound/rest"
	repository "hexa-design/internal/adapters/outbound/database"
	"hexa-design/internal/domain/services"
	"hexa-design/internal/ports"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts_integration(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expected := ""
		productRepo := repository.NewProductRepositoryMock()
		productRepo.On("GetProducts").Return(ports.Product{
			ID:       271,
			Name:     "Product271",
			Quantity: 99,
		}, nil)
		productSvc := services.NewProductService(productRepo)
		productHandler := handler.NewProductHandler(productSvc)

		app := fiber.New()
		app.Get("/", productHandler.GetProducts)

		req := httptest.NewRequest("GET", "/", nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, expected, body)
		}
	})
}
