package main

import (
	h "github.com/6ixBit/v2-Personal-Website/handlers"
	"github.com/robfig/cron"
)

// ScheduleTasks to be executed such as emailing
func ScheduleTasks() {
	c := cron.New()
	c.AddFunc("@every 4h", h.UpdateProjects)
	c.Start()
}
