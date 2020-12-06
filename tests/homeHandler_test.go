package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	h "github.com/6ixBit/v2-Personal-Website/handlers"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cv", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Response recorder to monitor test request.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HomeHandler)
	handler.ServeHTTP(rr, req) 	 					// Write to response recorder the test request

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong response code: %d", rr.Code)
	}

	if header := rr.Header().Get("Content-type"); header != "applicaton/pdf" {
		t.Errorf("Wrong header type set received from server.")
	}
}