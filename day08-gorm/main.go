package main

import (
	"golang-days/day08-gorm/controllers"
	"golang-days/day08-gorm/database"
)

func main() {
	database.StartDB()

	// controllers.CreateUser("azwarrasyid@gmail.com")
	// controllers.GetUserById(1)
	// controllers.UpdateUserById(1, "afintachtiars@gmail.com")
	// controllers.GetAllUsers()
	// controllers.CreateProduct(1, "Acer", "Yoga")
	controllers.GetAllUsers()
	controllers.GetAllUsersWithProducts()
}