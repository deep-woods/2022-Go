package main
import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 505
	var message string = "I'm going back"
	var isMonkey bool = true
	var amount float32 = 7.45

	// find type using %v
	fmt.Printf("var grades = %v, is of type %T\n", grades, grades)
	fmt.Printf("var message = %v, is of type %T\n", message, message)
	fmt.Printf("var isMonkey = %v, is of type %T\n", isMonkey, isMonkey)
	fmt.Printf("var amount = %v, is of type %T\n", amount, amount)

	// find type using reflect.TypeOf()
	fmt.Printf("Type: %v \n", reflect.TypeOf(grades))
	fmt.Printf("Type: %v \n", reflect.TypeOf(message))
	fmt.Printf("Type: %v \n", reflect.TypeOf(isMonkey))
	fmt.Printf("Type: %v \n", reflect.TypeOf(amount))
	fmt.Printf("Type: %v \n", reflect.TypeOf(45.7))
}