package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 10

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("value of a: %v\n", a)
	fmt.Printf("value of &a: %v\n", &a)
	fmt.Printf("value of b: %v\n", b)
	fmt.Printf("value of &b: %v\n\n", &b)

	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("value of p2: %v\n", p2)
	fmt.Printf("value of p3: %v\n\n", p3)

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n\n", p2 == p3)

	fmt.Printf("&a == p1: %v\n", &a == p1)
	fmt.Printf("&a == p2: %v\n", &a == p2)
	fmt.Printf("&a == p3: %v\n", &a == p3)
	fmt.Printf("&b == p3: %v\n\n", &b == p3)

	b = 20
	fmt.Println("b = 20")

	fmt.Printf("a == *p1: %v\n", a == *p1)
	fmt.Printf("a == *p2: %v\n", a == *p2)
	fmt.Printf("a == *p3: %v\n", a == *p3)
	fmt.Printf("b == *p3: %v\n", b == *p3)
}

/*
value of a: 10
value of &a: 0xc000112000
value of b: 10
value of &b: 0xc000112008

value of p1: 0xc000112000
value of p2: 0xc000112000
value of p3: 0xc000112008

p1 == p2: true
p2 == p3: false   <--------- SEE BELOW.

Although the values in the memory address are the same,
they are respectively stored in different memory locations.
Hence, the comparison results to be false.

&a == p1: true
&a == p2: true
&a == p3: false
&b == p3: true

b = 20
a == *p1: true
a == *p2: true
a == *p3: false
b == *p3: true

*/
