package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// MongoDB connection URI
	uri := "mongodb+srv://malikrtamboli:763YmElFWeCO3TAV@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"

	// Create a context with a timeout for connecting to the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	// Disconnect the client when the function is done
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Access the MongoDB collection
	collection := client.Database("testdb").Collection("users")

	// Filter to specify which document to delete
	filter := bson.M{"name": "Malik Tamboli"}

	// Perform the delete operation
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	if result.DeletedCount > 0 {
		fmt.Printf("Successfully deleted %d document(s)\n", result.DeletedCount)
	} else {
		fmt.Println("No documents matched the filter")
	}
}
