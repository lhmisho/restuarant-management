package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	OrderId    string    `json:"order_id"`
	OrderDate   time.Time `json:"order_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TableId   string    `json:"table_id"`
}
