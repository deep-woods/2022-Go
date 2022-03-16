package main
import "fmt"


// A type 
type Person struct {
	name string
	age int
}

// and its methods
func (p Person) GetName() string {
	return p.name
}

func (p Person) GetAge() int {
	return p.age
}

// Create a student struct including all the methods above. 
type Student struct {
	Person			// TYPE EMBEDDING
	studentId int
	major string
}

func (s Student) GetStudentId() int {
	return s.studentId
}

func (s Student) GetMajor() string {
	return s.major
}


func main() {
	// var s1 = Student{ name: "Mariachi", age: 23, studentId: 160892, major: "Computer Science"} // cannot use promoted field Person.name in struct literal of type Student?
	var s1 = Student{ studentId: 160892, major: "Computer Science"} // {{ 0} 160892 Computer Science}
	fmt.Println(s1)	
}