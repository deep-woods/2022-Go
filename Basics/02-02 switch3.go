package main

import "fmt"

func main() {
	var a, b int = 7, 45
	switch {
		case a*b == 505: 
			fmt.Println("equal to 505")
		case a*b < 505:
			fmt.Println("less than 505")
		default:
			fmt.Println("greater than 505")
	}
}