package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/mproni/users/internal/database"
	"github.com/mproni/users/internal/models"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUserByID(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
	}

	var user models.User
	stmt := "SELECT id, name, age, description FROM users WHERE ID = ?"
	row := database.DB.QueryRow(stmt, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Description)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	stmt := "INSERT INTO users (name, age, description) VALUES (?, ?, ?)"
	res, err := database.DB.Exec(stmt, newUser.Name,
		newUser.Age, newUser.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	newUser.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO
}
