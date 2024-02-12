package controllers

import (
	"golang-days/day08-task/database"
	"golang-days/day08-task/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoreOrder(c *gin.Context) {
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

func IndexOrder(c *gin.Context) {
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

func ShowOrder(c *gin.Context) {
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

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var updatedOrder models.Order

	c.ShouldBindJSON(&updatedOrder)
	
	db := database.GetDB()

	err := db.Where("order_id = ?", id).Delete(&models.Item{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete items before update order",
			"error_detail": err,
		})
		return
	}

	temp, _ := strconv.Atoi(id)
	updatedOrder.OrderID = uint(temp)

	err = db.Model(&updatedOrder).Where("order_id", id).Updates(updatedOrder).Error

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

func DestroyOrder(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDB()

	errOrder := db.Delete(&models.Order{}, id).Error
	if errOrder != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete order",
			"error_detail": errOrder,
		})
		return
	}

	errItem := db.Where("order_id = ?", id).Delete(&models.Item{}).Error
	if errItem != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete item",
			"error_detail": errItem,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "successfully delete order and its item",
	})
}