package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validator:"required, min=2, max=20"`
	Last_name     *string            `json:"last_name" validator:"required, min=2, max=20"`
	Password      *string            `json:"password" validator:"required, min=6"`
	Email         *string            `json:"email" validator:"email, required"`
	Phone         *string            `json:"phone" validator:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validator:"required, eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
