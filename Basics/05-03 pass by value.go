package main
import "fmt"

func modify(s string) string {	// changes the value of the input string
	s = "world"
	return s
}

// When the value is passed into a function, 
// the original value of the variable remains unchanged. 
func main() { 
	s := "hello"
	fmt.Printf("%v ", s)  
	s_mod := modify(s)
	fmt.Println(s_mod) // outputs: "hello world"
	fmt.Println(s)	   // outputs the original value "hello"
}