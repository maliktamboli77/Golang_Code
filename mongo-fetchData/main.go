package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://malikrtamboli:763YmElFWeCO3TAV@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust")
	clientOptions.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func fetchCollectionData(client *mongo.Client, dbName, collectionName string) ([]bson.M, error) {
	// Get the collection
	collection := client.Database(dbName).Collection(collectionName)

	//Prepare a filter
	filter := bson.D{}

	// Fetch documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	client, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Fetch data from a collection
	results, err := fetchCollectionData(client, "Bucket", "Personinfo")
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, result := range results {
		fmt.Println(result)
	}
}
