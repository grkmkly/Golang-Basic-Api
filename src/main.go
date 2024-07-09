package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

var users []user

// This function get users inside users slice
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := convertToStruct()
	w.Write(users)
}

// This function get user to given the id
func getUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	userIDstr := parts[2]

	// userIDstr to Integer
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		panic(err)
	}

	// ID match usersID
	for _, value := range users {
		if value.ID == userID {
			w.Header().Set("Content-Type", "application/json")
			jsonValue, err := json.Marshal(value)
			if err != nil {
				panic(err)
			}
			w.Write(jsonValue)
		}
	}
}

// This function finds maxID of the users
func getMaxId(users []user) int {
	maxIndex := len(users) - 1
	return users[maxIndex].ID
}

// This function writes from form value to json file and adds form value to  users struct
func postUser(w http.ResponseWriter, r *http.Request) {
	x, err := os.ReadFile("htmlFiles/postIndex.html")
	if err != nil {
		panic(err)
	}
	w.Write(x)
	if r.Method == http.MethodPost {
		usernameU := r.FormValue("username")
		emailU := r.FormValue("email")
		maxID := getMaxId(users)
		userU := user{
			ID:       maxID + 1,
			Username: usernameU,
			Email:    emailU,
		}
		users = append(users, userU)
		convertToJson()
	}
}

// This function writes users to json file
func convertToJson() {
	var text string = "[\n"
	maxIndex := getMaxId(users) - 1
	for index, value := range users {
		if index == maxIndex {
			text += fmt.Sprintf("\t{\"id\" : %v , \"name\" : \"%v\" ,\"email\" : \"%v\" }\n]", value.ID, value.Username, value.Email)
			break
		}
		text += fmt.Sprintf("\t{\"id\" : %v , \"name\" : \"%v\" ,\"email\" : \"%v\" },\n", value.ID, value.Username, value.Email)
	}
	os.WriteFile("example.json", []byte(text), 0755)
}

// This function converts inside json file to Struct
func convertToStruct() []byte {
	x, err := os.ReadFile("example.json")
	if err != nil {
		log.Fatal(err)
	}
	er := json.Unmarshal(x, &users)
	if er != nil {
		log.Fatal(err)
	}
	return x
}

// Main
func main() {

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/{id}", getUser)
	http.HandleFunc("/", postUser)
	http.ListenAndServe(":8000", nil)
}
