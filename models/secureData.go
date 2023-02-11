package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SecureDataModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Content   string             `bson:"content" validate:"required"`
	User      string             `bson:"user,omitempty"`
	HashCode  string             `bson:"hashCode,omitempty"`
	FinderKey string             `bson:"finderKey,omitempty"`
	Url       string             `bson:"url,omitempty"`
	// Protected bool               `bson:"protected"`
}

// type LoginBody struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// }
