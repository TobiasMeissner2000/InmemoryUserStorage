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
	fmt.Println("...Start server")
	fmt.Printf("Port : %s \n", port)
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

	err = t.Execute(w, items)

	if err != nil {
		fmt.Printf("Failed to start tamplate with error: %s ", err)
		return
	}
}

func AddUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("-- Add user --")

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u := new(user.User)
	err := json.NewDecoder(req.Body).Decode(u)

	if err != nil {
		fmt.Printf("Decode failed with err %s \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = req.Body.Close()
	u.SetId()
	userStorage[u.Id] = *u

	u.Print()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.MarshalIndent(u, "", "\t")

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)

	fmt.Printf("-- user with id %d added -- \n", u.Id)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("-- Delete user --")

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
	fmt.Println("-- Get user --")

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
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

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("-- write user with id %d \n--", id)
	w.Write(jsonResp)
	return
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("--  Get all user --")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.MarshalIndent(userStorage, "", "\t")
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
