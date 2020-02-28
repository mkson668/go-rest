package resthandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetServerRequest(t *testing.T) {

	// no query parameters for request.GET
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder for our server responses
	responseRec := httptest.NewRecorder()

	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)

	// feed server get request
	s.ServeHTTP(responseRec, req)

	// get the response code from the response recorder
	status := responseRec.Code

	// check if produce correct output
	if status != http.StatusOK {
		t.Errorf("Error code is incorrect, wanted %v, but got %v", http.StatusOK, status)
	}
}

func TestPostServerRequest(t *testing.T) {

	// no query parameters for request.GET
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder for our server responses
	responseRec := httptest.NewRecorder()

	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)

	// feed server get request
	s.ServeHTTP(responseRec, req)

	// get the response code from the response recorder
	status := responseRec.Code

	// check if produce correct output
	if status != http.StatusCreated {
		t.Errorf("Error code is incorrect, wanted %v, but got %v", http.StatusCreated, status)
	}
}

func TestPutServerRequest(t *testing.T) {

	// no query parameters for request.GET
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder for our server responses
	responseRec := httptest.NewRecorder()

	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)

	// feed server get request
	s.ServeHTTP(responseRec, req)

	// get the response code from the response recorder
	status := responseRec.Code

	// check if produce correct output
	if status != http.StatusAccepted {
		t.Errorf("Error code is incorrect, wanted %v, but got %v", http.StatusAccepted, status)
	}
}

func TestDeleteServerRequest(t *testing.T) {

	// no query parameters for request.GET
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder for our server responses
	responseRec := httptest.NewRecorder()

	s := server{}
	// set URL pattern to listen for
	http.Handle("/", s)

	// feed server get request
	s.ServeHTTP(responseRec, req)

	// get the response code from the response recorder
	status := responseRec.Code

	// check if produce correct output
	if status != http.StatusOK {
		t.Errorf("Error code is incorrect, wanted %v, but got %v", http.StatusOK, status)
	}
}
