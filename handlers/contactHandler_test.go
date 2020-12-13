package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
)

func TestContactHandler(t *testing.T) {
	testData := ContactForm{
		Name: "Tester",
		Email: "test@test.com",
		Subject: "Just testing things",
		Message: "Post request test to contact endpoint",
	}
	// Marshal struct to bytes then plug bytes into a reader object
	// so post req can be interpreted
	testBody,_ := json.Marshal(testData)
	testReader := bytes.NewReader(testBody)

	req, err := http.NewRequest("POST", "/api/contact", testReader)
	req.Header.Set("Content-type", "application/json")
	if err != nil { t.Fatal(err) }

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ContactHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong response code: %d", rr.Code)
	}
}
