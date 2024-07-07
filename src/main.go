package main

import (
	"encoding/json"
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

// GET USERS INSIDE STRUCT
func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	users, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(users)
}

// GET USERS GIVEN ID
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

// FIND MAX ID
func getMaxId(users []user) int {
	maxIndex := len(users) - 1
	return users[maxIndex].ID
}

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
	}
}
func main() {

	user1 := user{
		ID:       1,
		Username: "grkmkly35",
		Email:    "kolaygorkem@icloud.com",
	}
	user2 := user{
		ID:       2,
		Username: "Instagram",
		Email:    "instagram@gmail.com",
	}
	users = append(users, user1, user2)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/{id}", getUser)
	http.HandleFunc("/", postUser)
	http.ListenAndServe(":8000", nil)

}
