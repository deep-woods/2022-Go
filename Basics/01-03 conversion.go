package main
import (
	"fmt"
	"strconv"
)

// Itoa(): integer to string
// Atoi(): string to integer

func main() {
	// normal conversion	
	var f float64 = 505.54589 
	var i int = int(f)     // float64 is converted to integer. 
	fmt.Printf("%v\n", i)  // prints 505

	// integer to string using strconv package
	var i2 int = 505
	var s string = strconv.Itoa(i2) // convert int to string
	fmt.Printf("%q\n", s) 			// prints "505"

	// string to integer
	fmt.Printf("%v of type %T\n", s, s)
	var i3, e = strconv.Atoi(s)
	fmt.Printf("%v of type %T\n", i3, i3, e)
}

