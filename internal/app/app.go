package app

import (
	"net/http"

	"github.com/mproni/users/internal/database"
	"github.com/mproni/users/internal/handlers"
)

func Start() {
	db := database.InitDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.Users)
	mux.HandleFunc("/users/", handlers.User)

}
