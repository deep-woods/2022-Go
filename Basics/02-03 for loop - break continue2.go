package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i ==  3 {
			continue	// compare the result of "break" and "continue". 
		}				// in case of continue, it skips the remaining operation and jumps to i == 4.
		fmt.Println(i)
	}
}
