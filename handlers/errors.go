package handlers

import (
	"errors"
	"os"
	"log"
)

func init() {
	file, err := os.Create("errorLogs.txt")

	if err != nil {
		log.Println("File Error: Failed to create file to log errors to")
	}

	log.SetOutput(file)
}

var (
	// ErrFailedToParseJSON failed to parse request
	ErrFailedToParseJSON = errors.New("Failed to parse incoming JSON")
)
