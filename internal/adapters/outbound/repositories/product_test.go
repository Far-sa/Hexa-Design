package repository_test

import (
	repository "hexa-design/internal/adapters/outbound/repositories"
	"hexa-design/internal/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {

	db := setup(t)

	expected := model.Product{
		ID:       1,
		Name:     "",
		Quantity: 99,
	}

	t.Run("product not found", func(t *testing.T) {
		product := repository.NewProductRepositoryDb(db)
		_, err := product.GetProducts()
		assert.Nil(t, err)

	})
	t.Run("success", func(t *testing.T) {
		productDB := repository.NewProductRepositoryDb(db)
		product, err := productDB.GetProducts()
		assert.Nil(t, err)
		assert.Equal(t, expected, product)

	})
}
