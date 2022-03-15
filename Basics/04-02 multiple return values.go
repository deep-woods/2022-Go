package main
import "fmt"

// MULTIPLE RETURN VALUES
// Syntax 1
func operate1(a int, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}

// Syntax 2
func operate2(a int, b int) (mul int, div int) {
    mul = a * b     // equals symbol is gone.
    div = a / b    // same here. 
    return
}

func main() {
    sum, difference := operate1(100, 20)
    mul, div        := operate2(100, 20)
    fmt.Println(sum, difference)  // 120 80
    fmt.Println(mul, div)         // 2000 5
}
