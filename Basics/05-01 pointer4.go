package main
import "fmt"

func main() {
	var y int
	var ptr *int = &y

	*ptr = 500
	fmt.Println(ptr, *ptr)	// 0xc000016088 500

	*ptr += 5
	fmt.Println(ptr, *ptr)	// 0xc000016088 500
}