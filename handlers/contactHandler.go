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

var (
	incomingForm ContactForm
)

// ContactHandler is responsible for handling the contact route
func ContactHandler(w http.ResponseWriter, r *http.Request) {	
	if r.Method != "POST" {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Err405)

		return
	}

	// Parse request and write to contactForm struct
	responseInBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(responseInBytes, &incomingForm); err != nil { 
		log.Println(ErrFailedToParseJSON) 

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Err400)
		return
	} 
	
	// Temporarily Write back POST req to client
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(incomingForm); err != nil {
		log.Println("Failed to return response to client")
	}
} 