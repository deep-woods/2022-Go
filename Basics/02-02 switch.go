package main

import "fmt"

func main() {
	var i int = 10
	switch i {
		case -5: 
			fmt.Println("-5")
			fallthrough		
		case 10: 
			fmt.Println("10")
			fallthrough			// keeps falling thorugh and executing the following code block regardless of case match. 
		case 17: 
			fmt.Println("17")
			fallthrough
		default:
			fmt.Println("default")			
	}
} 