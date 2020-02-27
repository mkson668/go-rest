package main

import (
	"log"
	"net/http"
)

type server struct{}

// this is called composition in go, think of it as adding a function to our server struct
// we must define this fnction for our server to work
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// https://golang.org/pkg/net/http/#ResponseWriter
	// this will signify that request is successful
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	byteArr := []byte(`{"message" : "server reacheddd"}`)
	w.Write(byteArr)
}

func main() {
	// create server struct
	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)
	// start server and log info
	log.Fatal(http.ListenAndServe(":8080", nil))
}
