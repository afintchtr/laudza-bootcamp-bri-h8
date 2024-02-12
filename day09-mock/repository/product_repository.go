package repository

import "golang-days/day09-mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}