package routers

import (
	"golang-days/day08-swagger/controllers"
	_ "golang-days/day08-swagger/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Cars API
// @version 1.0
// @description This is what ive been practiced in Golang class Bootcamp BRI Golang
// @termsOfService http://swagger.io/terms/
// @contact.name Afin Tachtiar
// @contact.email afintchtr@gmail.com
// @license.name AfinT
// @license.url https://github.com/afintchtr
// @host localhost:3030
// @BasePath /
func StartServer() *gin.Engine {
	r := gin.Default()
	
	// Create
	r.POST("/cars", controllers.CreateCar)
	// Update
	r.PATCH("/cars/:id", controllers.UpdateCarById)
	// Get All
	r.GET("/cars", controllers.GetAllCars)
	// Get One
	r.GET("/cars/:id", controllers.GetCarById)
	// Delete
	r.DELETE("/cars/:id", controllers.DeleteCarById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}