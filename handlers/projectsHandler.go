package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	repos []Projects
	c = cache.New(cache.NoExpiration, 10*time.Minute)
)

// Projects is a placeholder for response from GitHubs API
type Projects struct {
	Name 			string `json: "name"`
	Svn_URL 		string `json: "svn_url"`
	Description 	string `json: "description"`
	Language 		string `json: "language"`
}

// ProjectsHandler is responsbile for dealing with projects route
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(Err405)
		
		return
	}
	
	// Fetch current projects from cache and return response.
	if projects, exists := c.Get("projects"); exists {

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&projects)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(Err500)
	}
}

// Fetches projects from GitHub API and writes response to addrss of repos global
func FetchProjects() {
	url, _ := os.LookupEnv("GITHUB_URL")

	res, err := http.Get(url)
	if err != nil { log.Println(ErrSendingRequest)}
	defer res.Body.Close()

	// Read response from response body
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {	log.Println(ErrFailedToReadFromReader) }

	err = json.Unmarshal(response, &repos)
	setProjectsInCache(repos)			

	if err != nil { 
		log.Println(ErrFailedToParseJSON)
	}
}

func setProjectsInCache(r []Projects) {
	c.Set("projects", &r, cache.NoExpiration)
	_, exists := c.Get("projects")

	if exists == false {
		log.Println("Failed to write to projects to cache")	
	}
}

// UpdateProjects will fetch repos, write result to in memory struct, then set that value to cache.
// Cache read is executed to respond to user request for repos.
func UpdateProjects() {
	FetchProjects()
}