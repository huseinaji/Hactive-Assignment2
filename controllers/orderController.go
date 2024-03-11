package controllers

import (
	"go_restapi_assignment2/config"
	"go_restapi_assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TesServer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "server is running",
	})
}

func CreateOrders(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Create(&order).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "data created",
	})
}

func GetOrders(ctx *gin.Context) {
	var order models.Order
	err := config.DB.Preload("Items").Find(&order).Error
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": order,
	})
}
