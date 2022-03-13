package main
import "fmt"

func main() {
	const name = "Hermione Granger"
	const is_muggle = false
	const age = 12

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n\n", age, age)

	name = "Harry Potter"

	fmt.Printf("%v is %v years old.", name, age)

}
