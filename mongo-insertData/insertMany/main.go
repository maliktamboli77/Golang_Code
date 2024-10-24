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
	// Set the MongoDB URI
	uri := "mongodb+srv://malikrtamboli:763YmElFWeCO3TAV@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB using mongo.Connect() directly
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Define the database and collection
	collection := client.Database("testdb").Collection("users")

	// Define multiple documents to insert
	documents := []interface{}{
		bson.D{{"name", "Malik Tamboli"}, {"age", 27}},
		bson.D{{"name", "Akash Jadhav"}, {"age", 27}},
		bson.D{{"name", "Bob Marley"}, {"age", 25}},
	}

	// Insert multiple documents
	insertManyResult, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	// Output inserted documents IDs
	fmt.Println("Inserted document IDs: ", insertManyResult.InsertedIDs)

	// Disconnect from MongoDB
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
