package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingRoutes.POST("/orders/create", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/update", controllers.UpdateOrder())
}