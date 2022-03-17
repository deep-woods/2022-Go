package main

import (
	"fmt"
)

func main() {)
	go runSomeLoop(14)
	fmt.Println("Hello, world")
}

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Prnting:", i)
	}
}