package main

import (
	"log"
	"net/http"

	"github.com/6ixBit/v2-Personal-Website/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	startServer()
}

func startServer() {
	log.Println("Server up and running on port 8080..")
	http.ListenAndServe(":8080", nil)
}
