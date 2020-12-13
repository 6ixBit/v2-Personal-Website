package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/projects", nil)
	if err != nil { t.Fatal(err) }

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProjectsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong response code: %d", rr.Code)
	}
}