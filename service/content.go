package service

import (
	"context"
	"log"
	"time"

	"github.com/raza11409652/securepass/database"
	"github.com/raza11409652/securepass/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var contentCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "contents")

func InsertNewContent(c models.SecureDataModel) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := contentCollection.InsertOne(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}
