package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// User define user struct
type User struct {
	ID    int
	Name  string
	Score int
}

// Users represent result comming from DB
var Users = []User{
	{
		ID:    0,
		Name:  "Eslam",
		Score: 1,
	},
	User{
		ID:    1,
		Name:  "Mahmoud",
		Score: 2,
	},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/", index)

	fmt.Println("starting server")

	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling / req")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"Hello world :)"}`)
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		get(w, req)
	} else if req.Method == "POST" {
		post(w, req)
	} else {
		fmt.Println("Handling invalid /users req")
		errorHandler(w, req, http.StatusMethodNotAllowed, fmt.Errorf("Invalid Method"))
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling GET req")
	// http://localhost:8090/users?id=0
	query := req.URL.Query()
	id := query.Get("id")
	var result []byte
	var err error

	// if no id
	if id == "" {
		result, err = json.Marshal(Users)
	} else {
		// if we have id convert it from string to int
		idInt, err := strconv.Atoi(id)
		if err == nil {
			result, err = json.Marshal(Users[idInt])
		}
	}

	// if we had any error return status 500 and error
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

	// set header return data
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(result))
}

func post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling POST req")
	var u User
	defer req.Body.Close()

	// read req body
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

	// decode the body
	err = json.Unmarshal(b, &u)
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

	// Users represent result comming from DB
	// so we act as if we saved it into DB
	Users = append(Users, u)

	w.WriteHeader(http.StatusCreated)
}

func errorHandler(w http.ResponseWriter, req *http.Request, status int, err error) {
	w.WriteHeader(status)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{error:%v}`, err.Error())
}
