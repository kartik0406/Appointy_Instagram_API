package handler
import (
	"context"
	"encoding/json"
	"instagram_api/controller"
	"instagram_api/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetPost(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter.  Using bson.M{} as it is not necessary to sort data
	filter := bson.M{"_id": id}
	err := postCollection.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		controller.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(post)
}
