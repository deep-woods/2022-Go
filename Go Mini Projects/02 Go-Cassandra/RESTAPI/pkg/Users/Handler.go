package Users

import (
	database "RESTAPI/pkg/Database"
	"encoding/json"
	"net/http"
)

var (
	user User
)

// make the functions private.
func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me!")
}

// Handler == controller (the entrypoint within the rest service)
// map it to 'user' model, add it to DB, then return a result confirmation.
func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)
	addUser(&user)

	println("just added user to db...")
	json.NewEncoder(w).Encode("You just added user1 to cassandra DB")
}

func addUser(user *User) {
	query := `INSERT INTO users(FirstName, LastName, Eamil) VALUES (?, ?, ?)`
	database.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}

/*
func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just sent a post request.")
}
*/
