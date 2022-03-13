package main

import "fmt"

func main() {

	// Syntax 
	for i:=1; i<=5; i++ {
		fmt.Println(i*i)
	}

	// Same effect as above
	j := 1
	for j <=5 {
		fmt.Println(j*j)
		j += 1
	}
}
