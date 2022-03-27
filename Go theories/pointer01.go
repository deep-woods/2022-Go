package main

import (
	"fmt"
)

func main() {
	var a int = 500
	var p *int

	p = &a
	fmt.Printf("value of p: %p\n", p)
	fmt.Printf("value in the memory address which p points to:  %d\n", *p)

	*p = 100
	fmt.Printf("value of a: %d\n", a)
}

/*

value of p: 0xc0000b8000
value in the memory address which p points to:  500
value of a: 100    <----- value of a changed just by manipulating p.

*/
