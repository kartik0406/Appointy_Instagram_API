package main

import (
	
	"net/http"
	"net/http/httptest"
	"testing"
	route "instagram_api/handler"
)

func TestGetPost(t *testing.T) {
	// data :={"id":1,"caption":"hello this is my first post","imageurl":"https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg","userid":"667187481661"}
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(route.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	
	if rr.Body.String() == ""{
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "data")
	}
}

func TestGetUser(t *testing.T) {
	// data :={"id":1,"name":"Krish","email":"krish405@gmail.com","password":"1234"}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(route.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	
	if rr.Body.String() == ""{
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "data")
	}
}
