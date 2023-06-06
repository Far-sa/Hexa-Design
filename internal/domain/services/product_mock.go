package services

import (
	"hexa-design/internal/ports"

	"github.com/stretchr/testify/mock"
)


type productServiceMock struct{
	mock.Mock
}
func NewProductServiceMock()*productServiceMock{
	return &productServiceMock{}
}

func (m *productServiceMock)GetProducts()([]ports.Product,error){
	args := m.Called()
	return args.Get(0).([]ports.Product),args.Error(1)
}