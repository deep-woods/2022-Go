package Router

import (
	database "RESTAPI/pkg/Database"
	"RESTAPI/pkg/Users"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

/*

Understand ROUTING.
- Routing is the process of selecting a path for traffic in a network or between or across multiple networks.

ROUTING in web development.
Routing or router in web development is a mechanism
where HTTP requests are routed to the code that handles them.
To put simply, in the Router you determine what should happen when a user visits a certain page.
*/

var port = ":7199"

func StartServer() *chi.Mux {

	database.SetupDBConnection()

	r := chi.NewRouter()

	// Mount the routes defined in /Users/Router.go
	r.Mount("/api/users", Users.UsersRoutes())
	fmt.Println("Server is listening on port" + port)
	log.Fatal(http.ListenAndServe(port, r))
	return r
}

/* REFERENCES

- https://divpusher.com/glossary/routing/#:~:text=Routing%20or%20router%20in%20web,user%20visits%20a%20certain%20page.
*/
