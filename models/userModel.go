package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required, min=2"`
	Last_name     *string            `json:"last_name" validate:"required, min=2"`
	Password      *string            `json:"password" validate:"required, min=4"`
	Email         *string            `json:"email" validate:"email, required"`
	Avater        *string            `json:"avater"`
	Phone         *string            `json:"phone" validate:"required,min=11"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Uupdated_at   time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
