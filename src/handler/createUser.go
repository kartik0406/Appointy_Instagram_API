package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"instagram_api/controller"
	"instagram_api/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User

	cur, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		controller.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var user models.User

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our User model.
	data := []byte("hello")
	hash := md5.Sum(data)
	user.Password = hex.EncodeToString(hash[:])
	result, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		controller.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
