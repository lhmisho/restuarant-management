package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"restuarant-management/routes"
)

func main(){
	fmt.Println("I'm working")
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.FoodRoutes(router)
	routes.UserRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + port)
}