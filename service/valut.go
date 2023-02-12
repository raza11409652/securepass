package service

import (
	"context"
	"log"
	"time"

	"github.com/raza11409652/securepass/database"
	"github.com/raza11409652/securepass/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var vaultCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "secure_vaults")

func InsertNewVault(c models.SecureVault) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := vaultCollection.InsertOne(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}
