package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id"`
	Order_id    string             `json:"order_id"`
	Order_date  time.Time          `json:"order_date"`
	Created_at  time.Time          `json:"created_at"`
	Uupdated_at time.Time          `json:"updated_at"`
	Table_id    string             `json:"table_id"`
}
