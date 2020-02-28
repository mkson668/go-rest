package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type server struct{}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	byteArr := []byte(`{"message" : "Get Resource"}`)
	w.Write(byteArr)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	byteArr := []byte(`{"message" : "Created Resource"}`)
	w.Write(byteArr)
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	byteArr := []byte(`{"message" : "Updated Resource"}`)
	w.Write(byteArr)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	byteArr := [[]byte(`{"message" : "Deleted Resource"}`)
	w.Write(byteArr)
}

func notValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	byteArr := [[]byte(`{"message" : "Not a valid method"}`)
	w.Write(byteArr)
}


func handleRequests(w http.ResponseWriter, r *http.Request) {
	// https://golang.org/pkg/net/http/#ResponseWriter
	// this will signify that request is successful
	w.Header().Set("Content-Type", "application/json")
	// we should set the reponse type for each type of HTTP request method
	var byteArr []byte
	switch requestType := r.Method; requestType {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Println("Get some resource")
		byteArr := []byte(`{"message" : "Updated Resource"}`)
		
	case "POST":
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Create some resource")
		byteArr := []byte(`{"message" : "Created Resource"}`)
		
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		fmt.Println("Update some resource")
		byteArr := []byte(`{"message" : "Updated Resource"}`)
		
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		fmt.Println("Deleted some resource")
		byteArr := []byte(`{"message" : "Deleted Resource"}`)
		
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Not a valid method")
		byteArr := []byte(`{"message" : "Not a valid method"}`)
		
	}
	w.Write(byteArr)
}

func main() {
	r := mux.NewRouter()
	// set URL pattern to listen for
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notValid)
	// start server and log info
	log.Fatal(http.ListenAndServe(":8080", r))
}
