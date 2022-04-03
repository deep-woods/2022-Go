package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var PORT string = ":3030"

// fmt.Sprintf("%s", os.Getenv("PORT"))

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	fmt.Fprintf(w, "Hello Foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	nameQuery := fmt.Sprintf("Hello %v", name)
	fmt.Fprintf(w, nameQuery)
}

func NewHttpHandler() http.Handler {

	mux := http.NewServeMux()

	// route 1 root
	mux.HandleFunc("/", IndexHandler)

	mux.HandleFunc("/bar", barHandler) // route 2

	mux.HandleFunc("/bar?name", barHandler)

	mux.Handle("/foo", &fooHandler{}) // route 3

	http.ListenAndServe(PORT, mux)
	return mux
}
