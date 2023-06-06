package repository_test

import (
	"errors"
	repository "hexa-design/internal/adapters/outbound/database"
	"testing"
)

func TestGetProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	cases := []testCase{
		{test: "err not found", name: "", expectedErr: errors.New("")},
		{test: "success", name: "", expectedErr: nil},
	}

	for _, c := range cases {
		t.Run(c.test, func(t *testing.T) {
			productRepo := repository.NewProductRepositoryMock()
			_, err := productRepo.GetProducts()
			if !errors.Is(err, c.expectedErr) {
				t.Error("", c.expectedErr, err)
			}
		})
	}
}
