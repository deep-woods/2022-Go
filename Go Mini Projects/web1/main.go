package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct{}

// ServeHTTP is an interface; fooHandler class' function.
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bar!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Foo!")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Forest"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

type jsonHandler struct{}

// This function handles json data.
func (j *jsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)                           // for user data
	err := json.NewDecoder(r.Body).Decode(user) // decode json-format data.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request:", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

https://pkg.go.dev/net/http
*/

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

func main() {
	// create a router.
	mux := http.NewServeMux()

	// HandleFunc maps a handler function to a specific route.
	// r holds the information requested by the client.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/About", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, this is Forest.")
	})

	mux.HandleFunc("/Contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Reach me at 🌳🌳🌳")
	})

	// Use Handle() when passing in an instance (&fooHandler in this case).
	mux.Handle("/Foo", &fooHandler{})
	mux.HandleFunc("/Bar", barHandler)
	/* Compare Syntax
	http.Handle("/route", &struct{})
	http.HandleFunc("/route", func)
	*/

	mux.Handle("/json", &jsonHandler{})
	mux.HandleFunc("/name", nameHandler)

	http.ListenAndServe(":3030", mux)
}

/*
func http.ListenAndServe(addr string, handler http.Handler) error
http.ListenAndServe on pkg.go.dev

ListenAndServe listens on the TCP network address addr
and then calls Serve with handler to handle requests on incoming connections.
Accepted connections are configured to enable TCP keep-alives.
The handler is typically nil, in which case the DefaultServeMux is used.

ListenAndServe always returns a non-nil error.
*/

// fhttps://www.youtube.com/watch?v=4Oml8mbBXgo&list=PLy-g2fnSzUTDALoERcKDniql16SAaQYHF
