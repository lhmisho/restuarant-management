package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	Quantity    *string            `json:"quantity"`
	UnitPrice   *float64           `json:"unit_price"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	OrderItemId string             `json:"order_item_id"`
	OrderId     string             `json:"order_id"`
	FoodId      *string            `json:"food_id"`
}
