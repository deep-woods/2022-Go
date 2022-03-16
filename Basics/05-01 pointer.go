package main
import "fmt"

func main() {
	i := 505
	fmt.Printf("%T %v \n", &i, &i)		  // *int 0xc00013a000
	fmt.Printf("%T %v \n", *(&i), *(&i))  // int 505
}