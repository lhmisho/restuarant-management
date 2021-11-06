package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUsers())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
}