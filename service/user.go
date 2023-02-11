package service

import (
	"context"
	"log"
	"time"

	"github.com/raza11409652/securepass/database"
	"github.com/raza11409652/securepass/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "users")

func GetUsers() *mongo.Cursor {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	crr, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return crr

}

func InsertNewUser(user models.UserModel) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}

func GetUserCount(filter interface{}) int64 {
	// fmt.Print(filter)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal((err))
	}
	defer cancel()
	return result
}

func GetUser(filter interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result := userCollection.FindOne(ctx, filter)
	defer cancel()
	return result
}
