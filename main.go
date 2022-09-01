package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"test/user"
	"time"
)

const (
	port = ":8080"
)

//struct httphandler{ userstorage }
//handler funktionen als methoden definieren

//mutex
//unterschied methode, funtion, pseudo construktor
//rest auf struct get, del,...
//handling map and slices

var userStore = user.NewUserStore()

func main() {
	fmt.Println("...Start server")
	fmt.Printf("Port : %s \n", port)
	router := createRouter()
	startServer(router)
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", printReadme)
	router.HandleFunc("/add", AddUser)
	router.HandleFunc("/delete", DeleteUser)
	router.HandleFunc("/get", GetUser)
	router.HandleFunc("/getall", GetAllUser)

	return router
}

func startServer(serveMux *http.ServeMux) {
	server := http.Server{
		Addr:        port,
		Handler:     serveMux,
		ReadTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
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
	userStore.Add(u)

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

	fmt.Printf("-- user with id %d added -- \n", u.ID)
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

	err = userStore.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
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
	id, err := strconv.Atoi(myId) //convert string to int
	if err != nil {
		fmt.Println("Failed with error: " + err.Error())
		return
	}

	fmt.Printf("Id: %d \n", id)
	u, err := userStore.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, err.Error())
		//w.Write([]byte(err.Error()))
		return
	}

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
	jsonResp, err := json.MarshalIndent(userStore.GetAll(), "", "\t")
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s \n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Printf("-- write users --")
	w.Write(jsonResp)
	return
}
