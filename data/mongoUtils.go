package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	godotenv.Load(".env")
	OpenConnection()
}

// MongoDatastore - MongoDB access object
type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
}

// DataStore - Holds mongo database & client
var DataStore *MongoDatastore

// OpenConnection opens a MongoDB connection
func OpenConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:" + os.Getenv("PASSWORD") + "@" + os.Getenv("DB_NAME") + "-i52u9.mongodb.net")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		DataStore = new(MongoDatastore)
		DataStore.Session = client
		DataStore.db = client.Database(os.Getenv("DB_NAME"))
		fmt.Println("Connected to MongoDB!")
	}

	return client
}

// GetCollection - returns collection
func GetCollection() *mongo.Collection {
	collection := DataStore.db.Collection(os.Getenv("COLLECTION_NAME"))
	return collection
}
