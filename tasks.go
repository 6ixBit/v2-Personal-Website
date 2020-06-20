// Background tasks to be done over a period of time
package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func tt() {
	fmt.Println("Gimme mioney rn")
}

func EmailRequestLogs() {
	gocron.Every(1).Second().Do(tt)
}