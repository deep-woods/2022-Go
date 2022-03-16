package main
import "fmt"

func modify(s *string) {
	*s = "world"
}

func main() {
	a:= "hello"
	fmt.Println(a) // "hello"
	modify(&a)	   // actual value modified
	fmt.Println(a) // "world"
}