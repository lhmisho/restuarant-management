package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *string            `json:"quantity"`
	Unit_price    *float64           `json:"unit_price"`
	Created_at    time.Time          `json:"created_at"`
	Uupdated_at   time.Time          `json:"updated_at"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id"`
	Food_id       *string            `json:"food_id"`
}
