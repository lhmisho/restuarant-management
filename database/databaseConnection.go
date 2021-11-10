package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DDinstance() *mongo.Client  {
	//MongoDB := "mongodb+srv://admin:admin@cluster0.xs68z.mongodb.net/restaurant_management?retryWrites=true&w=majority"
	MongoDB := "mongodb://mongoadmin:mongoadmin@192.168.0.136:27017"
	fmt.Println(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil{
		log.Println(err)
	}
	fmt.Println("Connected to mongodb")
	return client

}
var Client * mongo.Client = DDinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("restaurant_management").Collection(collectionName)
	return collection
}