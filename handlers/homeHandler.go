package handlers

import (
	"net/http"
	"io"
	"os"
	"log"
	"encoding/json"
)

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{http.StatusMethodNotAllowed, "HTTP Method not supported"})
		return
	}

	// Send PDF 
	file, err := os.Open("./HamzaCV.pdf")
	if err != nil { 
		log.Println(err) 
		w.WriteHeader(http.StatusInternalServerError)
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
