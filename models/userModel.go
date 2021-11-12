package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required, min=2"`
	LastName     *string            `json:"last_name" validate:"required, min=2"`
	Password     *string            `json:"password" validate:"required, min=4"`
	Email        *string            `json:"email" validate:"email, required"`
	Avatar       *string            `json:"avatar"`
	Phone        *string            `json:"phone" validate:"required,min=11"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}
