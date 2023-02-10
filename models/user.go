package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
	FirstName   *string            `bson:"firstName"`
	LastName    *string            `bson:"lastName"`
	Email       *string            `bson:"email" validate:"email required"`
	LastLoginAt *bool              `bson:"lastLoginAt"`
}
