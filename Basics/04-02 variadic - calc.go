package main
import "fmt"

// Variadic functions - takes in a varying number of arguments.
func calcSquare(numbers []int) []int {
        squares := []int{}
        for _, v := range numbers {
                squares = append(squares, v*v)
        }
        return squares
}


func calcSquare_bool(numbers []int) ([]int, bool) { // <--- notice the variadic input and multiple different return types. 
    squares := []int{}
    for _, v := range numbers {
            squares = append(squares, v*v)
    }
    return squares, true

}

func main() {
        nums := [3]int{10, 20, 15}
        fmt.Println(calcSquare(nums[:]))  // [100 400 225]
        fmt.Println(calcSquare_bool(nums[:]))  // [100 400 225]
}
