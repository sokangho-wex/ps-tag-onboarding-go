package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(uri string) *MongoClient {
	ctx := context.Background()
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return &MongoClient{client: client}
}

func (mc *MongoClient) NewMongoUserDB() *mongo.Database {
	return mc.client.Database("user")
}

func (mc *MongoClient) DisconnectMongoDB() {
	if err := mc.client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
