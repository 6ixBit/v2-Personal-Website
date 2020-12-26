package handlers

import (
	"net/http"
)

// Set headers for writer to allow for cross origin requests
func AllowCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}