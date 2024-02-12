package service

import (
	"golang-days/day09-mock/entity"
	"golang-days/day09-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProuctService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	prouct, err := productService.GetOneProduct("1")

	assert.Nil(t, prouct)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "error response has to be 'Product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product {
		Id: "2",
		Name: "Laptop",
	}
	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Laptop'")
	assert.Equal(t, &product, result, "result has to be data with id '2'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []*entity.Product {
		{
			Id: "1",
			Name: "Laptop",
		},
		{
			Id: "2",
			Name: "Handphone",
		},
	}
	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, len(products), len(result), "result length has to be 2")
	for i := 0; i < len(products); i++ {
		assert.Equal(t, products[i].Id, result[i].Id, "result has to be '1 or 2'")
		assert.Equal(t, products[i].Name, result[i].Name, "result has to be 'Laptop or Handphone'")
		assert.Equal(t, products[i], result[i], "result has to be data with id '1' or '2'")
	}
}