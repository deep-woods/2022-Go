package main

import (
	"app/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", app.NewHttpHandler())
}

/*
func main() {
	http.ListenAndServe(":3333", app.NewHttpHandler)   <---- Missing parenthesis.
}

cannot use app.NewHttpHandler (value of type func() http.Handler) as http.Handler value in argument to http.ListenAndServe: func() http.Handler does not implement http.Handler (missing method ServeHTTP)

https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#InvalidIfaceAssign
*/
