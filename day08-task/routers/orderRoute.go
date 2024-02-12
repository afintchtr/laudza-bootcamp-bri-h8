package routers

import (
	"golang-days/day08-task/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.FetchAllOrderWithItems)
	r.GET("/orders/:id", controllers.FetchOrderWithItemsById)
	r.GET("/orders/items/", controllers.FetchAllItems)
	r.GET("/orders/items/:id", controllers.FetchItemById)
	r.PATCH("/orders/:id", controllers.UpdateOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)

	return r
}