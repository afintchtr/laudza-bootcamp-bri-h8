package routers

import (
	"golang-days/day08-task/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	
	r.POST("/orders", controllers.StoreOrder)
	r.GET("/orders", controllers.IndexOrder)
	r.GET("/orders/:id", controllers.ShowOrder)
	r.GET("/orders/items/", controllers.IndexItem)
	r.GET("/orders/items/:id", controllers.ShowItem)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.PUT("/orders/items/:id", controllers.UpdateItem)	
	r.DELETE("/orders/:id", controllers.DestroyOrder)
	r.DELETE("/orders/items/:id", controllers.DestroyItem)

	return r
}