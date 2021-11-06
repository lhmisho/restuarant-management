package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRoutes.POST("/menus/create", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/update", controllers.UpdateMenu())
}