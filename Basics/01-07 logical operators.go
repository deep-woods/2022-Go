package main

import "fmt"

func main() {
	var x int = 505
	fmt.Println( (x < 100) && (x <= 505))
	fmt.Println( (x < 3000) && (x < 0))

	var a bool = true
	result := 7 > 45
	fmt.Println(!(a || result))

	var b bool = false
	result := 7 > 45
	fmt.Println(!(b || result))

}