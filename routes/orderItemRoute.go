package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/order-items", controllers.GetOrderItems())
	incomingRoutes.GET("/order-items/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.POST("/order-items/create", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order-items/update", controllers.UpdateOrderItem())
}