package controllers

import (
	"fmt"
	"golang-days/day08-gorm/database"
	"golang-days/day08-gorm/models"
)

func CreateProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product {
		UserID: userId,
		Brand: brand,
		Name: name,
	}
	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating Product data", err.Error())
		return
	}

	fmt.Println("New Product data:", Product)
}