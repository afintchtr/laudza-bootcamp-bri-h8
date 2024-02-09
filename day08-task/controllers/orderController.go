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
	})
}