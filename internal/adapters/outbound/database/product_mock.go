package repository

import (
	"hexa-design/internal/ports"

	"github.com/stretchr/testify/mock"
)


type productRepositoryMock struct{
	mock.Mock
}


func NewProductRepositoryMock()*productRepositoryMock{
	return &productRepositoryMock{}
}

func (m *productRepositoryMock) GetProducts()([]ports.Product,error){
	args := m.Called()
	return args.Get(0).([]ports.Product),args.Error(1)
}