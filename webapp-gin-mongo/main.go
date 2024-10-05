package main

import (
	"context"
	"log"
	"net/http"

	//Gin framework package
	"github.com/gin-gonic/gin"

	//mongodb driver packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoDB Atlas connection string
const uri = "mongodb+srv://malikrtamboli:763YmElFWeCO3TAV@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"

// Global variable that will hold a reference to the mongodb client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMoviesByID)
	r.POST("/movies/aggregation", aggregateMovies)

	r.Run()
}

func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

func getMovies(c *gin.Context) {
	// Find movies
	cursor, err := mongoClient.Database("sample_mflix").Collection("movies").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map results
	var movies []bson.M
	if err = cursor.All(context.TODO(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movies
	c.JSON(http.StatusOK, movies)
}

// Get movies by ID
func getMoviesByID(c *gin.Context) {

	// Get movie ID from URL
	idStr := c.Param("id")

	// Convert id string to ObjectId
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find movie by ObjectId
	var movie bson.M
	err = mongoClient.Database("sample_mflix").Collection("movies").FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movie
	c.JSON(http.StatusOK, movie)
}

// Aggregate Movies
func aggregateMovies(c *gin.Context) {
	var pipeline interface{}
	if err := c.ShouldBindJSON(&pipeline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Run aggregations
	cursor, err := mongoClient.Database("sample_mflix").Collection("movies").Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Map results
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return result
	c.JSON(http.StatusOK, result)
}
