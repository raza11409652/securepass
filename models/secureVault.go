package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SecureVault struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	User           primitive.ObjectID `bson:"user,omitempty"`
	Team           primitive.ObjectID `bson:"team,omitempty"`
	Type           string             `bson:"type" validate:"oneof=PASSWORD SECURE_NOTE"`
	MasterPassword bool               `bson:"masterPassword,omitempty"`
	Url            string             `bson:"url,omitempty"`
	Username       string             `bson:"username,omitempty"`
	Password       string             `bson:"password,omitempty"`
	Name           string             `bson:"name" validate:"required"`
	Notes          string             `bson:"notes,omitempty" validate:"required"`
	FileAttached   string             `bson:"fileAttached,omitempty"`
	VaultKey       string             `bson:"vaultKey,omitempty" `
}
