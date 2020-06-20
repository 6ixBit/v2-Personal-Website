package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"time"
	"fmt"

	cache "github.com/patrickmn/go-cache"
	// "github.com/prprprus/scheduler"
)

var (
	repos []Projects
	c = cache.New(cache.NoExpiration, 10*time.Minute)
)

// Projects is a placeholder for response from GitHubs API
type Projects struct {
	Name 		string `json: "name"`
	CloneURL 	string `json: "clone_url"`
	Description string `json: "description"`
	Language 	string `json: "language"`
}

// ProjectsHandler is responsbile for dealing with projects route
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	fetchProjects()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

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
	setProjectsInCache(repos)				// TESTING Cache Read and Write

	if err != nil { 
		log.Println(ErrFailedToParseJSON)
	}
}

func setProjectsInCache(r []Projects) {
	c.Set("projects", &r, cache.NoExpiration)
	items, exists := c.Get("projects")

	// IF item is found.
	if exists == true {
		fmt.Println(items)
	}
}

