package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection string

const (
	ProductsCollection collection = "products"
	// UsersCollection collection = "users"
)

const (
	url      = "mongodb://localhost:27017"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	var clientInstance *mongo.Client
	var clientInstanceError error

	var mongoOnce sync.Once

	mongoOnce.Do(func() {
		// const client = new MongoClient("mongodb://localhost:27017")
		clientOptions := options.Client().ApplyURI(url)

		// await client.connect()
		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
