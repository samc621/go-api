package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User object
type User struct {
	ID        int      `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address object
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var users []User

// GetAllUsers returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		json.NewEncoder(w).Encode("There are no users.")
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUserByID returns a user by their unique ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No user could be found with that ID.")
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// UpdateUserByID updates a user by their unique ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	for index, item := range users {
		if strconv.Itoa(item.ID) == params["id"] {
			users[index].Firstname = user.Firstname
			users[index].Lastname = user.Lastname
			users[index].Address = user.Address

			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	json.NewEncoder(w).Encode("No user could be found with that ID.")
}

// DeleteUserByID deletes a user by their unique ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if strconv.Itoa(item.ID) == params["id"] {
			users = append(users[:index], users[index+1:]...)

			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode("No user could be found with that ID.")
}

func main() {
	router := mux.NewRouter()

	//RESTful Routes
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUserByID).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserByID).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
