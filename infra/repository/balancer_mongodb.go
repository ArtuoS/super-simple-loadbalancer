package repository

import (
	"context"

	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "loadbalancer"

type BalancerMongoDB struct {
	client *mongo.Client
}

func NewBalancerMongoDB(client *mongo.Client) *BalancerMongoDB {
	return &BalancerMongoDB{
		client: client,
	}
}

func (b *BalancerMongoDB) PushServer(context context.Context, filter primitive.D, server primitive.D) (*mongo.UpdateResult, error) {
	return b.client.Database(databaseName).Collection(database.Balancers).UpdateOne(context, filter, server)
}

func (b *BalancerMongoDB) Search(context context.Context, filter primitive.D) (*mongo.Cursor, error) {
	return b.client.Database(databaseName).Collection(database.Balancers).Find(context, filter)
}

func (b *BalancerMongoDB) Update(context context.Context, filter primitive.D, update primitive.D, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return b.client.Database(databaseName).Collection(database.Balancers).UpdateOne(context, filter, update, options)
}
