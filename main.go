package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"test/user"
)

const (
	port = ":8080"
)

var (
	userStorage = make(map[int]user.User)
)

func main() {
	fmt.Println("--> Starting Server")
	handlerequests()
}

func handlerequests() {
	http.HandleFunc("/", printReadme)
	http.HandleFunc("/add", AddUser)
	http.HandleFunc("/delete", DeleteUser)
	http.HandleFunc("/get", GetUser)
	http.HandleFunc("/getall", GetAllUser)

	http.ListenAndServe(port, nil)
}

func printReadme(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("homePage.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	items := struct {
		Welcome string
	}{
		Welcome: "Here is the readme for this small project",
	}

	t.Execute(w, items)
}

func AddUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("-- add user --")

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u := new(user.User)
	json.NewDecoder(req.Body).Decode(u)
	req.Body.Close()
	u.SetId()
	userStorage[u.Id] = *u
	fmt.Println(u)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(u)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)

	fmt.Printf("-- user with id %d added --", u.GetId())
	fmt.Println()
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("-- delete user --")

	if req.Method != http.MethodDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	myId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(myId)
	if err != nil {
		fmt.Println("Failed with error: " + err.Error())
		return
	}

	if !contains(id) {
		fmt.Printf("Id %d not found \n", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(userStorage, id)
	w.WriteHeader(http.StatusCreated)

	fmt.Printf("-- user with id %d deleted -- \n", id)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("-- get user --")

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	myId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(myId)
	if err != nil {
		fmt.Println("Failed with error: " + err.Error())
		return
	}

	if !contains(id) {
		fmt.Printf("Id %d not found \n", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Printf("Id: %d \n", id)
	u := userStorage[id]

	u.Print()

	if err != nil {
		fmt.Printf("User with id %d not found \n", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(u)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Printf("-- write user with id %d --", id)
	w.Write(jsonResp)
	return
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(userStorage)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Printf("-- write users --")
	w.Write(jsonResp)
	return
}

func contains(id int) bool {

	for key := range userStorage {
		if key == id {
			return true
		}
	}
	return false
}
