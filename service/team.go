package service

import (
	"context"
	"log"
	"time"

	"github.com/raza11409652/securepass/database"
	"github.com/raza11409652/securepass/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var teamCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "teams")

func InsertNewTeam(team models.TeamData) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := teamCollection.InsertOne(ctx, team)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return result
}

func GetTeams(id primitive.ObjectID) *mongo.Cursor {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	crr, err := teamCollection.Find(ctx, bson.M{"createdBy": id})
	if err != nil {
		log.Fatal(err)
	}

	return crr

}
