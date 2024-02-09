package controllers

import (
	"errors"
	"fmt"
	"golang-days/day08-gorm/database"
	"golang-days/day08-gorm/models"

	"gorm.io/gorm"
)

func CreateUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data:", User)
}

func GetUserById(id uint) {
	db := database.GetDB()
	
	User := models.User{}

	err := db.First(&User, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found:")
			return
		}
		fmt.Println("Error finding user:", err)
	}
	fmt.Printf("User data: %+v \n", User)
}

func GetAllUsers() {
	db := database.GetDB()

	Users := []models.User{}

	err := db.Find(&Users).Error

	if err != nil {
		fmt.Println("Error get all user:", err)
		return
	}
	for i, User := range Users {
		fmt.Printf("User data no. %v : %+v \n", i+1, User)
	}
}

func UpdateUserById(id uint, newEmail string) {
	db := database.GetDB()

	User := models.User{}

	err := db.Model(&User).Where("id = ?", id).Updates(models.User{
		Email: newEmail,
	}).Error

	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Printf("Update user %v email to: %+v \n", User.ID, User.Email)
}

/////////////////////////////

func GetAllUsersWithProducts() {
	db := database.GetDB()

	Users := []models.User{}

	err := db.Preload("Products").Find(&Users).Error

	if err != nil {
		fmt.Println("Error get all user with products:", err.Error())
		return
	}
	fmt.Println("User data with Products")
	for i, User := range Users {
		fmt.Printf("User data no. %v : %+v \n", i+1, User)
	}
	// fmt.Printf("%+v", Users)
}