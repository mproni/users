package app

import (
	"net/http"

	"github.com/mproni/users/internal/database"
	"github.com/mproni/users/internal/handlers"
)

func Start() {
	database.InitDB()
	defer database.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.Users)
	mux.HandleFunc("/users/", handlers.User)

	http.ListenAndServe(":8090", mux)
}
