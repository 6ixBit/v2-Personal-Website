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
	stub.SendEmail()
	w.WriteHeader(http.StatusOK)
} 

func prepareEmailConfig(sender, body string) e.Email {
	srvAddr, _ := os.LookupEnv("SMTP_SERVER")
	srvPort, _ := os.LookupEnv("SMTP_PORT")
	recipient, _ := os.LookupEnv("SMTP_RECIPIENT")

	server := srvAddr + ":" + srvPort

	stub := e.Email{
		server,
		recipient,
		sender,
		body,
	}

	return stub
}