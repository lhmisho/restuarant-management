package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.POST("/foods/create", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/update", controllers.UpdateFood())
}