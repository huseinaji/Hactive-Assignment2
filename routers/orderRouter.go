package routers

import (
	"go_restapi_assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	// inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/", controllers.TesServer)
	router.POST("/order", controllers.CreateOrders)
	router.GET("/order", controllers.GetOrders)

	return router
}
