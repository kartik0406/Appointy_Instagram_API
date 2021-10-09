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

var postCollection, userCollection = controller.ConnectDB()

func GetUser(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	
	filter := bson.M{"_id": id}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		controller.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
	return 

}
