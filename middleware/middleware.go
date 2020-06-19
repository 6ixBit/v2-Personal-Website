package middleware

import (
	"net"
	"log"
	"net/http"
	"os"
	"fmt"
	"time"
)

// LogRequests tracks all the requests that hit an endpoint
func LogRequests(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Create("requestLogs.txt")
		if err != nil { log.Println(err) }

		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		curTime := time.Now()
		t := fmt.Sprintf("%d:%d:%d", curTime.Hour(), curTime.Minute(), curTime.Second())

		logger := log.New(file, t, 0)
		logger.Printf(" - %s - [ %s ] - %s \n", r.Method, r.URL, host)

		// Execute handler
		next(w, r)
	}
}
