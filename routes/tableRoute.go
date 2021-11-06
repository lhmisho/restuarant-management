package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables/create", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/update", controllers.UpdateTable())
}