package main

import (
	"fmt"
)

// PANIC HANDLING

func main() {
	panicTest(true)
	fmt.Println("Hello, world!")
}

func panicTest(p bool) { // call 1
	defer checkPanic() // call 2
	if p {
		panic("Panic requested")
	}
}

func checkPanic() {
	if r := recover(); r != nil {
		fmt.Println("A panic was captured, message:", r)
	}
}

/*
panicTest				- "Panic requested"
	defer checkPanic	- "A panic was captured, message:"

A panic was captured, message: Panic requested
Hello, world!
*/
