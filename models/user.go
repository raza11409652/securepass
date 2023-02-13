package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	FirstName    string             `bson:"firstName" validate:"required" `
	LastName     string             `bson:"lastName" json:"lastName"`
	Email        string             `bson:"email" validate:"email,required" `
	LastLoginAt  bool               `bson:"lastLoginAt"`
	Password     string             `bson:"password" validate:"required,min=8,max=255"`
	ProfileImage string             `bson:"profileImage"`
	Role         string             `bson:"role" validate:"oneof=ADMIN USER" `
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterBody struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" `
}
