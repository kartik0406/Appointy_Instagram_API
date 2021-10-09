package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Collection,*mongo.Collection) {

	// Used MongoDB Cluster for the project with database name Instagram_api and collections: posts and user 
	clientOptions := options.Client().ApplyURI("mongodb+srv://kartik:Test123@cluster0.ui1ab.mongodb.net/Instragram_api?retryWrites=true&w=majority")

	// Connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected To The Database")

	postCollection := client.Database("Instagram_api").Collection("posts")
	userCollection := client.Database("Instagram_api").Collection("users")

	return postCollection,userCollection
}


type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}


func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
