package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetUser_StatusNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost/get?id=1001", nil)
	w := httptest.NewRecorder()

	GetUser(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected %d, actual %d", http.StatusNotFound, resp.StatusCode)
	}
}

func Test_GetUser_StatusMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost/get?id=1001", nil)
	w := httptest.NewRecorder()

	GetUser(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Expected %d, actual %d", http.StatusNotFound, resp.StatusCode)
	}
}

func Test_createRouter(t *testing.T) {
	router := createRouter()

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(server.URL + "/getall")
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Status is %d", res.StatusCode)
	}
}
