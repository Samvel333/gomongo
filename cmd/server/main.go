package main

import (
	"log"
	"net/http"

	"github.com/samvel333/gomongo/api"
	"github.com/samvel333/gomongo/pkg/mongodb"
)

func main() {
	// Initialize MongoDB connection
	mongodb.Init()

	// Initialize the router
	r := api.SetupRouter()

	// Start the server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
