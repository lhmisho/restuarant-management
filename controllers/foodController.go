package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"restuarant-management/database"
	"restuarant-management/models"
	"time"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		foodID := c.Param("food_id")

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodID}).Decode(&food)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error while fetching food by food_id"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func GetFoods()  gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(food)
		if validationErr != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.MenuId}).Decode(&menu)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		food.CreatedAt = time.Now()
		food.UpdatedAt = time.Now()
		food.ID = primitive.NewObjectID()
		food.FoodId = food.ID.Hex()
		food.Price = 100.00

		result, insertErr := foodCollection.InsertOne(ctx, food)
		defer cancel()
		if insertErr != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		}
		c.JSON(http.StatusOK, result)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}