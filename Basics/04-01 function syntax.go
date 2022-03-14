package main

import "fmt"

// Calculate cube
func returnCube(n int) int {
	cube := n*n*n
	return cube
}

func main() {
	fmt.Println(returnCube(419))
}

// learning in Kodekloud https://kodekloud.com/topic/lab-function-syntax/