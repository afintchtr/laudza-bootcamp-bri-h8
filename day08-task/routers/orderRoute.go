package routers

import (
	"golang-days/day08-task/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)

	return r
}