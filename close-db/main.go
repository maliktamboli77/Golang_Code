package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Scope struct {
	Project string
	Area    string
}

type Note struct {
	Title string
	Tags  []string
	Text  string
	Scope Scope
}

var mdbClient *mongo.Client

func main() {
	const serverAddr string = "127.0.0.1:8081"
	// TODO: Replace with your connection string
	const connStr string = "mongodb+srv://malikrtamboli:763YmElFWeCO3TAV@testclust.ucq1m.mongodb.net/?retryWrites=true&w=majority&appName=testClust"
	done := make(chan struct{})

	fmt.Println("Hola Caracola")

	ctxBg := context.Background()
	var err error
	mdbClient, err = mongo.Connect(ctxBg, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mdbClient.Disconnect(ctxBg); err != nil {
			panic(err)
		}
	}()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP Caracola"))
	})
	router.HandleFunc("/notes", createNote)

	server := http.Server{
		Addr:    serverAddr,
		Handler: router,
	}
	server.RegisterOnShutdown(func() {
		defer func() {
			done <- struct{}{}
		}()
		fmt.Println("Signal shutdown")
		time.Sleep(5 * time.Second)
	})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}
	}()
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(os.Stderr, "HTTP server error %v\n", err)
		close(done)
	}
	<-done
}

func createNote(w http.ResponseWriter, r *http.Request) {
	// Ensure that the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming request body into a Note struct
	var note Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	// Get a reference to the "Notes" collection from MongoDB
	notesCollection := mdbClient.Database("NoteKeeper").Collection("Notes")

	// Insert the note into the collection
	result, err := notesCollection.InsertOne(r.Context(), note)
	if err != nil {
		http.Error(w, "Failed to save note to database", http.StatusInternalServerError)
		return
	}

	// Respond back with the inserted note and ID
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"status":     "success",
		"insertedId": result.InsertedID,
		"note":       note,
	})
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201 Created
	w.Write(jsonResponse)
}
