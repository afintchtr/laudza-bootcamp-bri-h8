package service

import (
	"errors"
	"golang-days/day09-mock/entity"
	"golang-days/day09-mock/repository"
)

type ProuctService struct {
	Repository repository.ProductRepository
}

func (service ProuctService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}
	return product, nil
}

func (service ProuctService) GetAllProduct() ([]*entity.Product, error) {
	products := service.Repository.FindAll()
	if len(products) != 2 {
		return nil, errors.New("Products must contain two elements")
	}
	return products, nil
}