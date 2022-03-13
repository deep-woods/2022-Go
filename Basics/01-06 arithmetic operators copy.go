package main

import "fmt"

func main() {
	// And & Or operators 
	var t, f bool = true, false
	fmt.Println(t && f)
	fmt.Println(t || f)

	fmt.Println(!t)
	fmt.Println(!f)

	var a bool = false
	result := 10 > 50
	fmt.Println(!(a && result))

	var b bool = true
	result := 10 > 50
	fmt.Println(!(b || result))

	// Modulus operation with %
	var c, d int = 505, 7
	fmt.Println(c % d)
}