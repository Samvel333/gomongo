package api

import (
	"github.com/gorilla/mux"
	"github.com/samvel333/gomongo/pkg/handler"
)

// SetupRouter initializes the routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/items", handler.GetItems).Methods("GET")
	r.HandleFunc("/items", handler.CreateItem).Methods("POST")

	return r
}
