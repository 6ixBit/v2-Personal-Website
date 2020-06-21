package main

import (
	"log"

	"github.com/prprprus/scheduler"
	h "github.com/6ixBit/v2-Personal-Website/handlers"
 	// "github.com/6ixBit/v2-Personal-Website/emails"
)

// ScheduleTasks to be executed such as emailing
func ScheduleTasks() {
	s, err := scheduler.NewScheduler(1000)

	if err != nil { log.Println(err) }

	// Update projects in cache every 3 hours
	s.Every().Hour(3).Do(h.UpdateProjects)

	// Send emails
}
