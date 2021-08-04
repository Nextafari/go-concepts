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
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
}

// Defining mongoclient
// var client *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func hashUserPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Print(err)
	}
	return string(bytes)
}

func mongoConnector() (*mongo.Client, error) {
	// uri := "mongodb+srv://nexta:foobar@cluster0.h9grc.mongodb.net/zuriChat?retryWrites=true"
	uri := DbName("DbName")

	clientOptions := options.Client().ApplyURI(uri)

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

	return client, nil
}

func createUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	client, error := mongoConnector()

	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}
	// User with unhashed password
	fmt.Println(user, ": is the First user")
	// Accessing the password field in the user struct to has the password
	user.Password = hashUserPassword(user.Password)
	// user with hashed password
	fmt.Println(user, ": is the Second user")
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
	router := mux.NewRouter()
	router.HandleFunc("/create-user", createUserEndpoint).Methods("POST")
	http.ListenAndServe(":8080", router)
}
