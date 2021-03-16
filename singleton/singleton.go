package singleton

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func GetInstance() *mongo.Database {
	if database == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, _ := mongo.Connect(ctx, clientOptions)
		database = client.Database("test")
	}
	return database
}
