package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i * i)
		if i == 3 {
				continue  // compare the result of "break" and "continue". 
		}
	}
}
