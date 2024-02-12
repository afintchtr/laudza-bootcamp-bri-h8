package controllers

import (
	"golang-days/day08-task/database"
	"golang-days/day08-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexItem(c *gin.Context) {
	var items []models.Item

	db := database.GetDB()
	err := db.Find(&items).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show any item",
			"error_detail":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func ShowItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	db := database.GetDB()
	err := db.First(&item, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't show this item",
			"error_detail":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem models.Item

	c.ShouldBindJSON(&updatedItem)

	db := database.GetDB()

	err := db.Model(&updatedItem).Where("item_id", id).Updates(updatedItem).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't update item",
			"error_detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
			"updated_order": updatedItem,
			"message": "successfully updated item",
	})
}

func DestroyItem(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDB()

	err := db.Where("item_id = ?", id).Delete(&models.Item{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete item",
			"error_detail": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "successfully delete item",
	})
}