package handlers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
)

//ContactForm placeholder for incoming json form data
type ContactForm struct {
	Name 	string `json:"name"` 
	Email 	string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// ErrorResponse is a placeholder for json error messages
type ErrorResponse struct {
	Status 	int `json: "status"`
	Message string `json: "message"`
}

var (
	incomingForm ContactForm
)

var (
	err400 = ErrorResponse{http.StatusBadRequest, "The data received does not match the schema criteria"}
	err405 = ErrorResponse{http.StatusBadRequest, "That http method is not allowed for this endpoint."}
)

// ContactHandler is responsible for handling the contact route
func ContactHandler(w http.ResponseWriter, r *http.Request) {	
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(err405)

		return
	}

	// Parse request and write to contactForm struct
	responseInBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(responseInBytes, &incomingForm); err != nil { 
		log.Println(ErrFailedToParseJSON) 
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err400)

		return
	} 
	
	// Write succesful response to client
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(incomingForm); err != nil {
		log.Println("Failed to return response to client")
	}
} 