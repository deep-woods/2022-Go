package main

import (
	"fmt"
	"strconv"
)

// DEFER
// defer calls a specific code after the parent code has been executed.

func printEnding(message string) {
	fmt.Println(message)
}

func doSomething() {
	for j := 0; j <= 5; j++ {
		defer printEnding("doSomething() just ended " + strconv.Itoa(j))
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	doSomething()
}
