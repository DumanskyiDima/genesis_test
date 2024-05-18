package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection

func GetClient() *mongo.Client {
	uri := os.Getenv("DATABASE_URL")

	if client != nil {
		return client
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectioName string) *mongo.Collection {
	if userCollection != nil {
		return userCollection
	}
	userCollection := client.Database("ExchangeRateTracker").Collection(collectioName)
	return userCollection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
