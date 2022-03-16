package main
import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
/*	c := s1{x: 4}
	c1 := s2{x: 4}
	if c == c1 { //  invalid operation: c == c1 (mismatched types s1 and s2)
		fmt.Println("Two structs are equal.")
	} */

	o := s1{x: 505}
	o1 := s1{x: 45}
	o2 := s1{x: 505}

	if o == o1 {
		fmt.Println("o and o1 have equal values.")
	} else if o == o2 {
		fmt.Println("o and o2 have equal values.")
	}
}