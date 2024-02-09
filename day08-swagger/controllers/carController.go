package controllers

import (
	"errors"
	"golang-days/day08-swagger/database"
	"golang-days/day08-swagger/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCar godoc
// @Summary Create a new car
// @Description Create one car with name and price
// @Tag cars
// @Accept json
// @Produces json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/ [post]
func CreateCar(c *gin.Context) {
	var body struct {
		Name string
		Price uint
	}
	c.Bind(&body)
	Car := models.Car{
		Name: body.Name,
		Price: body.Price,
	}
	db := database.GetDB()
	err := db.Create(&Car).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't create car",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"car": Car,
	})
}


// GetCarById godoc
// @Summary Get car details
// @Description Get details of one car by its id
// @Tag cars
// @Accept json
// @Produces json
// @Param id path int true "ID of the car to be shown"
// @Success 200 {object} models.Car
// @Router /cars/{id} [get]
func GetCarById(c *gin.Context) {
	id := c.Param("id")
	var Car models.Car
	db := database.GetDB()
	err := db.First(&Car, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "car not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't show car",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"car": Car,
	})
}


// GetAllCars godoc
// @Summary Get all cars details
// @Description Get details of all cars
// @Tag cars
// @Accept json
// @Produces json
// @Success 200 {object} models.Car
// @Router /cars/ [get]
func GetAllCars(c *gin.Context) {
	var Cars []models.Car
	db := database.GetDB()
	err := db.Find(&Cars).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "car not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cars": Cars,
	})
}


// UpdateCarById godoc
// @Summary Update an existing car
// @Description Update an existing car with new name and price
// @Tag cars
// @Accept json
// @Produces json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCarById(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name string
		Price uint
	}
	c.Bind(&body)
	db := database.GetDB()	
	var UpdatedCar models.Car
	db.First(&UpdatedCar, id)
	err := db.Model(&UpdatedCar).Updates(models.Car{
		Name: body.Name,
		Price: body.Price,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "car can't updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cars": UpdatedCar,
	})
}


// DeleteCarById godoc
// @Summary Delete an existing car
// @Description Just delete an existing car
// @Tag cars
// @Accept json
// @Produces json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCarById(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	db.Delete(&models.Car{}, id)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "car already deleted",
	})
}
