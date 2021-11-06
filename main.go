package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
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

	router.Run(":" + port)
}