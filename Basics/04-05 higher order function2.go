package main
import "fmt"

// HIGHER ORDER FUNCTION
// Function that receives a function as an argument.
// Function that returns another function.

// WHY?
// reduces bugs
// increase code readability

func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Calculation completed")
}


// higher-order function
func getFunction(query int) func(r float64) float64 {
	
	query_to_func := map[int]func(r float64) float64 {
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

func main() {
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter\n:")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Result: ", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Result: ", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Result: ", calcDiameter(radius))
	} else {
		fmt.Println("Invalid query")
	}
}