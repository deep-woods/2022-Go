package Users

import (
	"encoding/json"
	"net/http"
)

var (
	user User
	s    = new(service) // used just for the receiver for Service.Adduser.
)

// make the functions private.
func GetHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me!")
}

// Handler == controller (the entrypoint within the rest service)
// map it to 'user' model, add it to DB, then return a result confirmation.
func PostHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)
	s.AddUser(&user)

	println("just added user to db...")
	json.NewEncoder(w).Encode("You just added user1 to cassandra DB")
}

/*
func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just sent a post request.")
}
*/

// CPoP (2020) Adding Service Layer To GoLang Rest API https://www.youtube.com/watch?v=VxFbLcIy8j8&list=PLle8fNcWU5av-TANvnRXqiM5BECBY-nvj&index=7&ab_channel=CPoP
