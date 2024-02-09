package database

import (
	"fmt"
	"golang-days/day08-task/env"
	"golang-days/day08-task/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.Host, env.User, env.Password, env.DbName, env.DbPort)

	env.Db, env.Err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if env.Err != nil {
		log.Fatal("Error connecting to database:", env.Err)
	}
	fmt.Println("Successfully connected to database")
	// env.Db.Debug().AutoMigrate(models.User{}, models.Product{})
	env.Db.AutoMigrate(models.Item{})
	env.Db.AutoMigrate(models.Order{})
}

func GetDB() *gorm.DB {
	return env.Db
}