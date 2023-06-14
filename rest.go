package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new user
	newUser.ID = len(users) + 1

	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	// Initialize some example data
	users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	// Define your API endpoints and their corresponding handlers
	http.HandleFunc(" ", getUsers)
	http.HandleFunc("/users/add", addUser)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8082", nil))
}
