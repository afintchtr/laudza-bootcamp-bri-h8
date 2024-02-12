package controllers

import (
	"golang-days/day08-task/database"
	"golang-days/day08-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var currentOrder models.Order

	c.ShouldBindJSON(&currentOrder)

	db := database.GetDB()
	err := db.Create(&currentOrder).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't create order",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"created_order": currentOrder,
			"message": "successfully created order",
	})
}

func FetchAllOrderWithItems(c *gin.Context) {
	var orders []models.Order

	db := database.GetDB()
	err := db.Preload("ListOfItem").Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show any order",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"orders": orders,
	})
}

func FetchOrderWithItemsById(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	db := database.GetDB()
	err := db.Preload("ListOfItem").First(&order, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show this order",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"order": order,
	})
}

func FetchAllItems(c *gin.Context) {
	var items []models.Item

	db := database.GetDB()
	err := db.Find(&items).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show any item",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"items": items,
	})
}

func FetchItemById(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	db := database.GetDB()
	err := db.First(&item, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show this item",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"item": item,
	})
}

func UpdateOrderById(c *gin.Context) {
	id := c.Param("id")

	var updatedOrder models.Order
	c.ShouldBindJSON(&updatedOrder)
	
	db := database.GetDB()
	err := db.Model(&updatedOrder).Where("order_id", id).Updates(updatedOrder).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't update order",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"updated_order": updatedOrder,
			"message": "successfully updated order",
	})
}

func DeleteOrderById(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDB()
	err := db.Delete(&models.Order{}, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete item",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"message": "successfully delete order",
	})
}