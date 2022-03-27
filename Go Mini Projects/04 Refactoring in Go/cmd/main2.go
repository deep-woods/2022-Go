package main

import (
	"RESTAPI/pkg/Router"
	"fmt"
)

// RUN A SERVER IN GO 02 - Refactor code into separate directories.

func main() {
	fmt.Println("Starting server...")
	Router.StartServer()
}
