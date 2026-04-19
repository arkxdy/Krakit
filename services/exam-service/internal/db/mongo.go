package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func connectMongo(uri string) (*MongoClient, error) {
	// 1. Prepare Client Options from URI
	clientOptions := options.Client().ApplyURI(uri)

	maxRetries := 30
	retryDelay := 2 * time.Second

	var client *mongo.Client
	var err error

	for i := 0; i < maxRetries; i++ {
		// In v2, Connect returns a client immediately
		client, err = mongo.Connect(clientOptions)

		if err == nil {
			// CRITICAL: Connect only initializes the client.
			// We must Ping to verify the server is actually reachable.
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			err = client.Ping(ctx, nil)
			cancel()

			if err == nil {
				log.Printf("Successfully connected to MongoDB after %d attempts", i+1)
				return &MongoClient{Client: client}, nil
			}

			// If Ping fails, cleanup the unusable client before retrying
			_ = client.Disconnect(context.Background())
		}

		log.Printf("Attempt %d/%d: MongoDB not ready... retrying in %v", i+1, maxRetries, retryDelay)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to connect to MongoDB after %d attempts: %w", maxRetries, err)
}

func (c *MongoClient) Close() {
	c.Client.Disconnect(context.Background())
}
