package main
import "fmt"

func main() {

	var i *int
	var s *string
	fmt.Println(i)	// <nil>
	fmt.Println(s)  // <nil>
	
	// initialise 1
	// var <ptr name> *<data type> = $<var name>
	j := 505
	var ptr_j *int = &j
	
	// initialise 2
	// var <ptr name> = &<var name>  (use another defined variable)
	k := 45
	var ptr_k = &k

	// initialise 3 - shorthand 
	// <ptr name> := &<var name>
	l := "7 hours"
	ptr_l := &l

	fmt.Println("ptr_j", ptr_j)
	fmt.Println("ptr_k", ptr_k)
	fmt.Println("ptr_l", ptr_l)
}