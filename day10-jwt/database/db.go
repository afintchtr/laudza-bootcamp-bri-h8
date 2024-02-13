package database

import (
	"fmt"
	"golang-days/day10-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Host     = "localhost"
	User     = "postgres"
	Password = "ravenclaw500"
	DbPort   = "5432"
	DbName   = "golang-bootcamp-5"
	Db       *gorm.DB
	Err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Host, User, Password, DbName, DbPort)

	Db, Err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if Err != nil {
		log.Fatal("Error connecting to database:", Err)
	}
	fmt.Println("Successfully connected to database")
	Db.AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return Db
}