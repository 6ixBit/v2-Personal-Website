package main

import (
	"github.com/go-chi/httprate"
	"fmt"
	"log"
	"net/http"
	"time"

	h "github.com/6ixBit/v2-Personal-Website/handlers"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
    "github.com/prometheus/client_golang/prometheus/promhttp"
	
)

func init() {
	loadEnvFile()
	go h.FetchProjects()
	go ScheduleTasks()
	go startPrometheus()
}

func main() {	
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(50 * time.Second)) // Set timeout for incoming requests
	r.Use(httprate.LimitByIP(30, 1*time.Minute))
	r.Use(middleware.Heartbeat("/api/status"))

	r.Get("/api/cv", 	   h.HomeHandler)
	r.Get("/api/projects", h.ProjectsHandler)
	r.Post("/api/contact", h.ContactHandler)
	
	startServer(r)
}

func startServer(r *chi.Mux) {
	fmt.Println("Web Server up and running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println("Env File Error: Failed to load .env file with environment variable")
	}
}

func startPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Monitoring server started on port 4444...")
	http.ListenAndServe(":4444", nil)
}