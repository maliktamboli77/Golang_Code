package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Address struct {
	Street string
	City   string
	State  string
}

type Student struct {
	FirstName string  `bson:"first_name, omitempty"`
	LastName  string  `bson:"last_name, omitempty"`
	Address   Address `bson:"inline"`
	Age       int
}

func main() {
	// Connection URI
	uri := "mongodb+srv://malikrtamboli:yourpass@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	clientOps := options.Client()
	clientOps.SetMaxPoolSize(5)
	defer client.Disconnect(context.TODO())

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not to connect to MongoDB: ", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Database create and insert some data
	coll := client.Database("db").Collection("students")

	address1 := Address{"1 Main Road", "Dighanchi", "MH"}
	student1 := Student{FirstName: "Malik", LastName: "Tamboli", Address: address1, Age: 26}

	insertRes, err := coll.InsertOne(context.TODO(), student1)
	if err != nil {
		log.Fatal("Insertion err", err)
	}
	// Print acknowledgement of inserted database:
	fmt.Printf("Inserted document with ID: %v\n", insertRes.InsertedID)

}
