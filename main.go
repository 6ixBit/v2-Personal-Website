package main

import (
	"log"
	"net/http"

	"github.com/6ixBit/v2-Personal-Website/handlers"
	"github.com/joho/godotenv"
)

func init() {
	loadEnvFile()
}

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/projects", handlers.ProjectsHandler)

	startServer()
}

func startServer() {
	log.Println("Server up and running on port 8080..")
	http.ListenAndServe(":8080", nil)
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println("Env File Error: Failed to load .env file with environment variable")
	}
}
