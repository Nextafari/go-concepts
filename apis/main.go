package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func main() {
	fmt.Println("Starting Application...")
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:5580"))
	router := mux.NewRouter()
	router.HandleFunc("/my-test", hello).Methods("Get")
	http.ListenAndServe(":8080", router)

}
