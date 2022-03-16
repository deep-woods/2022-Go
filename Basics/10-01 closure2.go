package main
import "fmt"

func incrementor() func() int {
    i := 0					// i is created outside the anonymous func.
    return func() int { 	// This anonym function still knows and can grab i regardless.
        i++
        return i
    }
}

func main () {
	next := incrementor()  // next is a function returned by incrementor
	fmt.Println(next())    // prints 1
	fmt.Println(next())    // prints 2
	fmt.Println(next())    // prints 3
}
