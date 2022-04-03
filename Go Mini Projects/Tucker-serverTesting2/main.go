package main

import (
	"net/http"
	"testenv/app"
)

func main() {
	http.ListenAndServe(app.PORT, app.NewHttpHandler())
}
