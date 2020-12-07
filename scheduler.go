package main

import (
	"log" 

	h "github.com/6ixBit/v2-Personal-Website/handlers"
	"github.com/prprprus/scheduler"
	"fmt"
)

// ScheduleTasks to be executed such as emailing
func ScheduleTasks() {
	s, err := scheduler.NewScheduler(1000)

	if err != nil {
		log.Println(err)
		return 
	}

	// Update projects in cache every 6 hours
	jobID := s.Every().Minute(2).Do(h.UpdateProjects)
	if jobID != "" { fmt.Println(jobID, "executed")}
}
