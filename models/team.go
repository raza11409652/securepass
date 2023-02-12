package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamData struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name" validate:"required,min=5,max=128"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	CreatedBy      primitive.ObjectID `bson:"createdBy"`
	MasterPassword string             `bson:"masterPassword,omitempty"`
}
