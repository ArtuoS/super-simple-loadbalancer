package database

import (
	"context"
	"log"

	"github.com/ArtuoS/super-simple-loadbalancer/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitializeMongoDBClient() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	Client = client
}
