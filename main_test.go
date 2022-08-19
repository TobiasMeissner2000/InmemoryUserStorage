package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"test/user"
	"testing"
)

func Test_Contains(t *testing.T) {

	tests := []struct {
		name        string
		userStorage map[int]user.User
		inputID     int
		want        bool
	}{
		{
			name: "Contains",
			userStorage: map[int]user.User{

				1234: {
					Id: 1234,
				},

				1235: {
					Id: 1235,
				},
			},
			inputID: 1234,
			want:    true,
		},
		{
			name: "Does not contains",
			userStorage: map[int]user.User{
				1234: {
					Id: 1234,
				},
				1235: {
					Id: 1235,
				},
			},
			inputID: 1236,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userStorage = tt.userStorage
			got := contains(tt.inputID)
			if got != tt.want {
				t.Fatalf("got : %t, want : %t", got, tt.want)
			}
		})
	}
}

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

func Test_handlerequests(t *testing.T) {
	handlerequests()

	res, err := http.Get("http://localhost:8080/getall")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.StatusCode)

	a := res.StatusCode
	fmt.Println(a)

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Status is %d", res.StatusCode)
	}
}
