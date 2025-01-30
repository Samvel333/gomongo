package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Init initializes the MongoDB connection
func Init() {
	uri := "mongodb://mongo:27017" // Docker service name for MongoDB

	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("MongoDB connected")
}

// GetCollection returns the MongoDB collection for a given name
func GetCollection(name string) *mongo.Collection {
	return client.Database("test").Collection(name)
}
