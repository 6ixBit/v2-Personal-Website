package handlers

import (
	"net/http"
)

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	b := []byte("On the hello world.")
	w.Write(b)

	
}
