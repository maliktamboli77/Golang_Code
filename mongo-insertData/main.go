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
	//Connection URI
	uri := "mongodb+srv://malikrtamboli:Password@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"

	//Set a timeout context for mongoDb operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Connect to mongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to mongoDB: ", err)
	}
	fmt.Println("Connected to mongoDB!")

	//Create new data and collection
	database := client.Database("Bucket")

	collection := database.Collection("Personinfo")

	sampleData := bson.D{
		{Key: "name", Value: "Jack Sparrow"},
		{Key: "age", Value: 40},
		{Key: "city", Value: "New York"},
	}

	insertResult, err := collection.InsertOne(ctx, sampleData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
}
