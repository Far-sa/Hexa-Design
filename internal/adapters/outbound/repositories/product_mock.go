package repository

import (
	"hexa-design/internal/domain/model"

	"github.com/stretchr/testify/mock"
)


type productRepositoryMock struct{
	mock.Mock
}


func NewProductRepositoryMock()*productRepositoryMock{
	return &productRepositoryMock{}
}

func (m *productRepositoryMock) GetProducts()([]model.Product,error){
	args := m.Called()
	return args.Get(0).([]model.Product),args.Error(1)
}