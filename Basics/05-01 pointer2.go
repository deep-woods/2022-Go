package main
import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	arr_add := &arr
	fmt.Printf("%T %v \n", arr_add, *arr_add)
}