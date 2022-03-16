package main
import "fmt"

func modify(s []int) {
	s[0] += 100
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	modify(slice)
	fmt.Println(slice) // [101 2 3 4 5 6 7 8 9] 
	// unlike string, slice's original value is modified
	// because slice itself is a reference to an underlying array.
}