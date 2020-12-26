package handlers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"os"

	e "github.com/6ixBit/v2-Personal-Website/emails"
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
	AllowCORS(w)
	
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
	
	stub := prepareEmailConfig(incomingForm.Email, incomingForm.Message)
	go stub.SendEmail()
	w.WriteHeader(http.StatusOK)
		// Write JSON on success
} 

func prepareEmailConfig(sender, body string) e.Email {
	recipient, _ := os.LookupEnv("SMTP_RECIPIENT")

	var receivers []string
	receivers = append(receivers, recipient)

	stub := e.Email{
		receivers,
		sender,
		body,
	}
	return stub
}