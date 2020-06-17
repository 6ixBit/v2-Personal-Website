package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"fmt"
)

var (
	repos []interface{}
)

// ProjectsHandler is responsbile for dealing with projects route
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	fetchProjects()

	for _, val := range repos {
		fmt.Println(val)	
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&repos)
}

// Fetches projects from GitHub API and writes response to addrss of repos global
func fetchProjects() {
	url, _ := os.LookupEnv("GITHUB_URL")

	res, err := http.Get(url)
	if err != nil { log.Println("Request Error: Request to fetch github repos failed")}
	defer res.Body.Close()

	// Read response from response body
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {	log.Println("Reader Error: Failed to read bytes from response body.") }

	err = json.Unmarshal(response, &repos)
	if err != nil { 
		log.Println("JSON Parsing Error: Failed to parse JSON response")
		log.Println(err)
	}
}