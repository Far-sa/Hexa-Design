package services

import (
	"github.com/stretchr/testify/mock"
	"hexa-design/internal/domain/model"
)

type productServiceMock struct {
	mock.Mock
}

func NewProductServiceMock() *productServiceMock {
	return &productServiceMock{}
}

func (m *productServiceMock) GetProducts() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}
