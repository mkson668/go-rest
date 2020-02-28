package resthandler

import (
	"fmt"
	"net/http"
)

type server struct{}

// this is called composition in go, think of it as adding a function to our server struct
// we must define this fnction for our server to work
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		w.Write(byteArr)
	case "POST":
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Create some resource")
		byteArr := []byte(`{"message" : "Created Resource"}`)
		w.Write(byteArr)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		fmt.Println("Update some resource")
		byteArr := []byte(`{"message" : "Updated Resource"}`)
		w.Write(byteArr)
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		fmt.Println("Deleted some resource")
		byteArr := []byte(`{"message" : "Deleted Resource"}`)
		w.Write(byteArr)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Not a valid method")
		byteArr := []byte(`{"message" : "Not a valid method"}`)
		w.Write(byteArr)
	}
	w.Write(byteArr)
}

/* func main() {
	// create server struct
	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)
	// start server and log info
	log.Fatal(http.ListenAndServe(":8080", nil))
} */
