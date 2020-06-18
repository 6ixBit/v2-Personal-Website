package handlers

import (
	"errors"
	"os"
	"log"
)

var (
	// ErrFailedToParseJSON failed to parse request
	ErrFailedToParseJSON = errors.New("Failed to parse incoming JSON")
	ErrFailedToReadFromReader = errors.New("Reader Error: Failed to read bytes from response body.")
	ErrSendingRequest = errors.New("Request Error: Request to fetch github repos failed")
)

// ErrorResponse is a placeholder for json error messages
type ErrorResponse struct {
	Status 	int `json: "status"`
	Message string `json: "message"`
}

func init() {
	file, err := os.Create("errorLogs.txt")
	if err != nil { log.Println("File Error: Failed to create file to log errors to") }

	log.SetOutput(file)
}
