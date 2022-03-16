package main
import "fmt"

func adder() func(int) int {
	sum := 0 			       // var sum is created outside the anonym func.
	return func(x int) int {   // but it still works regardless.
		sum += x
		return sum
	}
}

func main() {

	sumClosure := adder()
	fmt.Println(sumClosure)
	fmt.Println(sumClosure(1))
	fmt.Println(sumClosure(2))

}