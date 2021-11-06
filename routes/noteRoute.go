package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/notes", controllers.GetNotes())
	incomingRoutes.GET("/notes/:note_id", controllers.GetNote())
	incomingRoutes.POST("/notes/create", controllers.CreateNote())
	incomingRoutes.PATCH("/notes/update", controllers.UpdateNote())
}