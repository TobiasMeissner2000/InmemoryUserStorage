package main

import (
	"net/http"
	"net/http/httptest"
	"test/user"
	"testing"
)

func Test_Contains(t *testing.T) {

	userStorage = make(map[int]user.User)

	a := contains(1)
	if a {
		t.Fatalf("Expected %t, actual: %t", false, a)
	}
}

func Test_GetUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost/get?id=1001", nil)
	w := httptest.NewRecorder()

	GetUser(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected %d, actual %d", http.StatusNotFound, resp.StatusCode)
	}
}
