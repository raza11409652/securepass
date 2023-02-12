package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/raza11409652/securepass/database"
	"github.com/raza11409652/securepass/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var contentCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "contents")
var contentHistoryCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "content_history")

func InsertNewContent(c models.SecureDataModel) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := contentCollection.InsertOne(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}

func GetContentByFilter(f interface{}) *mongo.SingleResult {
	fmt.Print(f)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result := contentCollection.FindOne(ctx, f)
	defer cancel()
	return result

}

func InsertNewContentHistory(c models.SecureDataHistoryModel) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := contentHistoryCollection.InsertOne(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}
func GetContentHistoryCount(filter interface{}) int64 {
	// fmt.Print(filter)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := contentHistoryCollection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal((err))
	}
	defer cancel()
	return result
}
