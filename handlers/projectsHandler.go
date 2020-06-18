package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/json"
)

var (
	repos []Projects
)

// Projects is a placeholder for response from GitHubs API
type Projects struct {
	Name 		string `json: "name"`
	HTMLURL 	string `json: "html_url"`
	Description string 	`json: "description"`
	Language 	string `json: "language"`
}

// ProjectsHandler is responsbile for dealing with projects route
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	fetchProjects()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&repos)
}

// Fetches projects from GitHub API and writes response to addrss of repos global
func fetchProjects() {
	url, _ := os.LookupEnv("GITHUB_URL")

	res, err := http.Get(url)
	if err != nil { log.Println(ErrSendingRequest)}
	defer res.Body.Close()

	// Read response from response body
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {	log.Println(ErrFailedToReadFromReader) }

	err = json.Unmarshal(response, &repos)
	if err != nil { 
		log.Println(ErrFailedToParseJSON)
	}
}

