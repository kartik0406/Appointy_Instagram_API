// Created by Kartik Khanna 19BCE0531 BTECH CSE 2019-2023 
package main

// Imported userdefined packages and standard libarary packages to the project
import (
	"fmt"
	"instagram_api/controller"
	"github.com/gorilla/mux"
	 route "instagram_api/handler"
	"log"
	"net/http"
)

//Connection of mongoDB with controller
var postCollection, userCollection = controller.ConnectDB()


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Instagram API</h1>")
}
func main() {

	doppelgeanger := http.NewServeMux()
	
	
	//Routes
	doppelgeaenger := mux.NewRouter()
	doppelgeaenger.HandleFunc("/", index)
	doppelgeaenger.HandleFunc("/posts", route.InsertPost).Methods("GET")
	doppelgeaenger.HandleFunc("/users", route.InsertUser).Methods("GET")
	doppelgeaenger.HandleFunc("/users/{id}", route.GetUser).Methods("GET")
	doppelgeaenger.HandleFunc("/posts/{id}", route.GetPost).Methods("GET")
	doppelgeaenger.HandleFunc("/posts/users/{id}", route.GetUserPost).Methods("GET")
	doppelgeaenger.HandleFunc("/posts", route.CreatePost).Methods("POST")
	doppelgeaenger.HandleFunc("/users", route.CreateUser).Methods("POST")

	//Server set up
	fmt.Println("Server Started at port 9000")
	log.Fatal(http.ListenAndServe(":9000", mux))

}
// Created by Kartik Khanna 19BCE0531 BTECH CSE 2019-2023 
