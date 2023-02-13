package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SecureDataModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	Content        string             `bson:"content" validate:"required"`
	User           string             `bson:"user,omitempty"`
	HashCode       string             `bson:"hashCode,omitempty"`
	FinderKey      string             `bson:"finderKey,omitempty"`
	Url            string             `bson:"url,omitempty"`
	MaxViewAllowed int64              `bson:"maxViewAllowed,omitempty"`
	// Protected bool               `bson:"protected"`
}

type ContentBody struct {
	Content        string `json:"content" validate:"required"`
	MaxViewAllowed int64  `json:"maxViewAllowed,omitempty"`
	// Protected bool               `bson:"protected"`
}
type SecureDataHistoryModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	IpAddress string             `bson:"ipAddress,omitempty"`
	Device    string             `bson:"device"`
	Referrer  string             `bson:"referrer,omitempty"`
	Content   primitive.ObjectID `bson:"content,omitempty"`
}
