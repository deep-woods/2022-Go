package main
import "fmt"

const PI float64 = 3.14  // global constant

func main() {
	var radius float64 = 45
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of Circle:", area)	
}