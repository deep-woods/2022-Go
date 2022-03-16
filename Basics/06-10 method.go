package main
import "fmt"

// A type
type Person struct {
	name string
	age int
}

// Methods
func (p Person) GetName()string{
	return p.name
}

func (p Person) GetAge()int{
	return p.age
}

var alex = Person { 
	name: "Alex Turner",
	age: 34,
}

func main() {
	fmt.Println(alex.GetName)
	fmt.Println(alex.GetAge)
}