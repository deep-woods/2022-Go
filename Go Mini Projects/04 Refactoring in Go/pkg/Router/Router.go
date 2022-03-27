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

	r := database.SetupDBConnection()

	us := Users.NewService(r)

	router := chi.NewRouter()
	router.Mount("/api/users", Users.UsersRoutes(us)) // Mount the routes defined in /Users/Router.go
	fmt.Println("Server is listening on port" + port)
	fmt.Printf("Docker IP: %s\n", database.IP)
	log.Fatal(http.ListenAndServe(port, router))
	return router
}

/* REFERENCES

By default, when you create or run a container using docker create or docker run, it does not publish any of its ports to the outside world. To make a port available to services outside of Docker, or to Docker containers which are not connected to the container’s network, use the --publish or -p flag. This creates a firewall rule which maps a container port to a port on the Docker host to the outside world. Here are some examples.

- Published ports https://docs.docker.com/config/containers/container-networking/

- https://divpusher.com/glossary/routing/#:~:text=Routing%20or%20router%20in%20web,user%20visits%20a%20certain%20page.

*/
