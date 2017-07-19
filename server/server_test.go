package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var req *http.Request
var err error

func TestHandler(t *testing.T) {
	_, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/",
		nil,
	)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
}

func TestResponseStatusOK(t *testing.T) {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterHandlers)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", rec.Code)
	}
}
