package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	mux.HandleFunc("/About", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, this is Forest.")
	})

	mux.HandleFunc("/Contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Reach me at 🌳🌳🌳")
	})

	mux.Handle("/Foo", &fooHandler{})
	mux.HandleFunc("/Bar", barHandler)

	mux.Handle("/json", &jsonHandler{})
	mux.HandleFunc("/name", nameHandler)
	return mux
}

// ====================================

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Requset:", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World, welcome to index page.")
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
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

https://pkg.go.dev/net/http
*/
