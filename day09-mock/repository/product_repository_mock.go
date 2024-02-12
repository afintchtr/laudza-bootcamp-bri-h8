package repository

import (
	"golang-days/day09-mock/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)

	return &product
}	

func (repository *ProductRepositoryMock) FindAll() []*entity.Product {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}
	// products := []*entity.Product{}
	// for _, argument := range arguments {
	// 	products = append(products, argument.(*entity.Product))
	// }
	products := arguments.Get(0).([]*entity.Product)
	return products
}	