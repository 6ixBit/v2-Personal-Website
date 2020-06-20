package middleware

import (
	"net"
	"log"
	"net/http"
	"os"
	"fmt"
	"time"
)

var (
	file *os.File
)

func init() {
	f, err := os.Create("requestLogs.txt")
	if err != nil { log.Println(err) }
	file=f
}

// LogRequests tracks all the requests that hit an endpoint
func LogRequests(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		curTime := time.Now()
		t := fmt.Sprintf("%d:%d:%d", curTime.Hour(), curTime.Minute(), curTime.Second())
		
		output := fmt.Sprintf("%s - %s - [ %s ] - %s \n", t, r.Method, r.URL, host)
		file.WriteString(output)
		
		// Execute handler
		next(w, r)
	}
}
