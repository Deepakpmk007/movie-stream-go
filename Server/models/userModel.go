package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId          string        `bson:"user_id" json:"user_id"`
	FirstName       string        `bson:"first_name" validate:"required,min=2,max=100" json:"first_name"`
	LastName        string        `bson:"last_name" validate:"required,min=2,max=100" json:"last_name"`
	Email           string        `bson:"email" validate:"required,email" json:"email"`
	Password        string        `bson:"password" validate:"required,min=6" json:"password"`
	Role            string        `bson:"role" validate:"oneof=ADMIN USER" json:"role"`
	CreatedAt       time.Time     `bson:"created_at" json:"create_at"`
	UpdatedAt       time.Time     `bson:"updated_at" json:"updated_at"`
	Token           string        `bson:"token" json:"token"`
	RefreshToken    string        `bson:"refresh_token" json:"refresh_token"`
	FavouriteGenres []Genre       `bson:"favourite_genres" validate:"required,dive" json:"favourite_genres"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,email"`
}

type UserResponse struct {
	UserID          string  `json:"user_id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	FavouriteGenres []Genre `json:"favourit_genres"`
}
