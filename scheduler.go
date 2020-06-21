package main

import (
	"log"

	h "github.com/6ixBit/v2-Personal-Website/handlers"
	"github.com/prprprus/scheduler"
	// "github.com/6ixBit/v2-Personal-Website/emails"
)

// ScheduleTasks to be executed such as emailing
func ScheduleTasks() {
	s, err := scheduler.NewScheduler(1000)

	if err != nil {
		log.Println(err)
		return 
	}

	// Update projects in cache every 3 hours
	s.Every().Hour(3).Do(h.UpdateProjects)

	// Send emails
}
