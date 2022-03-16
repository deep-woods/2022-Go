package main
import "fmt"

// user-defined data type
// a structure that groups data elements together.

/* SYNTAX
 type <struct name> struct {
	
} */

type Student struct {
	name string
	id int
	major string
	year int
	credits int
	courses []string
}

func main() {

	var s Student
	fmt.Printf("%+v \n", s)  // {name: id:0 major: year:0 credits:0 courses:[]}

	// Initialise 1
	// <var name> := new(<struct name>)
	s2 := new(Student)
	fmt.Printf("%+v \n", s2) 

	/* Initialise 2 - shorthand assignment
	 <var name> := <struct name> {
		 <field name>: <value>,
		 ...
	 } */ 

	s3 := Student{
		name: "Alex Turner",
		id: 505745,
		major: "music",
		year: 4,
		credits: 100,
		courses: []string{"modern art history", "pop", "rock", "punk"},
	}

	fmt.Printf("%v \n", s3)
	fmt.Printf("%+v", s3)
}

/*
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names

	https://pkg.go.dev/fmt
*/