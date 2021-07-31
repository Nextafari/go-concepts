package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// Defining mongoclient
var client *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func hello(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	fmt.Fprint(response, "Hello World, Kipptyhipptyruppyt Hundlebundleshandle")
}

func createUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var user *User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}
	fmt.Println(err)
	fmt.Printf("I am the user: %v \n", *user)
	collection := client.Database("zuriChat").Collection("user")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(result)

}

func main() {
	fmt.Println("Starting Application...")
	// client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/zuriChat")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Checking the connection
	// err = client.Ping(context.Background(), readpref.Primary())
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
	fmt.Println("Connected to MongoDB!")

	router := mux.NewRouter()
	router.HandleFunc("/my-test", hello).Methods("GET")
	router.HandleFunc("/create_user", createUserEndpoint).Methods("POST")
	http.ListenAndServe(":8080", router)
}
