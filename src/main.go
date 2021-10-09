// Created by Kartik Khanna 19BCE0531 BTECH CSE 2019-2023 
package main

// Imported userdefined packages and standard libarary packages to the project
import (
	
	"fmt"
	"instagram_api/controller"
	 route "instagram_api/handler"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	
)

//Connection of mongoDB with controller
var postCollection, userCollection = controller.ConnectDB()


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Instagram API</h1>")
}
func main() {

	mux := mux.NewRouter()
	
	//Routes
	mux.HandleFunc("/", index)
	mux.HandleFunc("/posts", route.InsertPost).Methods("GET")
	mux.HandleFunc("/users", route.InsertUser).Methods("GET")
	mux.HandleFunc("/users/{id}", route.GetUser).Methods("GET")
	mux.HandleFunc("/posts/{id}", route.GetPost).Methods("GET")
	mux.HandleFunc("/posts/users/{id}", route.GetUserPost).Methods("GET")
	mux.HandleFunc("/posts", route.CreatePost).Methods("POST")
	mux.HandleFunc("/users", route.CreateUser).Methods("POST")

	//Server set up
	fmt.Println("Server Started at port 9000")
	log.Fatal(http.ListenAndServe(":9000", mux))

}
// Created by Kartik Khanna 19BCE0531 BTECH CSE 2019-2023 