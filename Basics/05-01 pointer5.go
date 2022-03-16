package main
import "fmt"

func main() {
	var x int = 505
	var ptr *int = &x   // ptr as memory address
	fmt.Println(x)		// x value 505
	fmt.Println(*ptr)	// ptr value 505
	fmt.Println(ptr)	// address 0xc000016088
}