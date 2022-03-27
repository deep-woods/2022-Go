package main

// package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// RUN A SERVER IN GO 01 - Basic web server setup.

// func main() {
func main1() {
	fmt.Println("Starting server...")

	// declare a router.
	r := chi.NewRouter()
	r.Get("/api/get_test", getHandler)
	r.Post("/api/post_test", postHandler) // this line won't work if you test on a web browser. Use postman instead.

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// w: what you are sending back to the client
// r: request (what the client sends you.)
func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me!")
	// Business logic (service)
	// - ex) service adds a user to db.
	// - actual functionality that your applications needs to do.
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just sent a post request.")
}
