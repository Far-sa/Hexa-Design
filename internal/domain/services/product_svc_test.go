package services_test

import (
	"errors"
	repository "hexa-design/internal/adapters/outbound/database"
	"hexa-design/internal/domain/services"
	"hexa-design/internal/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts_service(t *testing.T) {
	type testCase struct {
		id       int
		name     string
		quantity int
	}

	cases := []testCase{
		{id: 271, name: "Product271", quantity: 99},
		{id: 171, name: "Product171", quantity: 99},
		{id: 371, name: "Product371", quantity: 99},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			productRepo := repository.NewProductRepositoryMock()

			expectedProducts := []ports.Product{
				{ID: c.id, Name: c.name, Quantity: c.quantity},
			}
			productRepo.On("GetProducts").Return(expectedProducts, nil)
			productSvc := services.NewProductService(productRepo)
			result, err := productSvc.GetProducts()

			assert.NoError(t, err)
			assert.Equal(t, expectedProducts, result)

			productRepo.AssertCalled(t, "GetProducts")
		})
		t.Run("", func(t *testing.T) {
			//* Arrange
			productRepo := repository.NewProductRepositoryMock()
			productRepo.On("GetProducts").Return([]ports.Product{}, errors.New(""))

			productSvc := services.NewProductService(productRepo)
			//* Act
			_, err := productSvc.GetProducts()

			//* Assert
			assert.Equal(t, err, errors.New("repo not found"))
			productRepo.AssertNotCalled(t, "GetProducts")
		})
	}
}
