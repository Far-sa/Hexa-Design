package services_test

import (
	"errors"
	repository "hexa-design/internal/adapters/outbound/repositories"
	"hexa-design/internal/domain/model"
	"hexa-design/internal/domain/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts_service(t *testing.T) {
	type testCase struct {
		id       int
		name     string
		quantity int
		expected model.Product
	}

	cases := []testCase{
		{id: 271, name: "Product271", quantity: 99},
		{id: 171, name: "Product171", quantity: 99},
		{id: 371, name: "Product371", quantity: 99},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			//* Arrange
			// expectedProducts := []model.Product{
			// 	{ID: c.id, Name: c.name, Quantity: c.quantity},
			// }
			productRepo := repository.NewProductRepositoryMock()
			productRepo.On("GetProducts").Return(model.Product{
				ID:       1,
				Name:     c.name,
				Quantity: c.quantity,
			}, nil)

			productSvc := services.NewProductService(productRepo)

			//* Act
			result, err := productSvc.GetProducts()
			expected := c.expected

			//* Assert
			assert.NoError(t, err)
			assert.Equal(t, expected, result)

			productRepo.AssertCalled(t, "GetProducts")
		})
		
		t.Run("", func(t *testing.T) {
			//* Arrange
			productRepo := repository.NewProductRepositoryMock()
			productRepo.On("GetProducts").Return([]model.Product{}, errors.New(""))

			productSvc := services.NewProductService(productRepo)
			//* Act
			_, err := productSvc.GetProducts()

			//* Assert
			assert.Equal(t, err, errors.New("repo not found"))
			productRepo.AssertNotCalled(t, "GetProducts")
		})
	}
}
