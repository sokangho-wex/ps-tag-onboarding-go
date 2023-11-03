package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(uri string) *Client {
	ctx := context.Background()
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return &Client{client: client}
}

func (mc *Client) CreateDB(name string) *mongo.Database {
	return mc.client.Database(name)
}

func (mc *Client) DisconnectDB() {
	if err := mc.client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
