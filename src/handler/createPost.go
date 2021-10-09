package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"instagram_api/controller"
	"instagram_api/models"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	
	
)

func InsertPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []models.Post

	cur, err := postCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		controller.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var post models.Post

		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&post)

	// insert our Post model.
	currentTime := time.Now()
	post.Timestamp = currentTime.String()
	fmt.Println(post.Caption)
	result, err := postCollection.InsertOne(context.TODO(), post)

	if err != nil {
		controller.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
