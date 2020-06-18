package handlers

import (
	"net/http"
	"io"
	"os"
	"log"
)

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Send PDF 
	file, err := os.Open("./HamzaCV.pdf")
	if err != nil { 
		log.Println(err) 
		return
	}
	defer file.Close()

	// Set header for PDF
	w.Header().Set("Content-type", "application/pdf")

	// Stream bytes from file to response writer
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, file); err != nil {
		log.Println(err)
	}
}
