package main
import "fmt"

func modify(arr [3]int) {
	for i := range arr {
			arr[i] -= 5
	}
}

/*
func modify_slice(s []int) []int {
	for i := range s {
		s[i] -= 5
	}
	return s
} */

func main() {
	arr := [3]int{10, 20, 30}
	fmt.Println(arr)  // [10 20 30]
	modify(arr)		  // original array NOT modified
	fmt.Println(arr)  // [10 20 30]

//	slice := modify_slice(arr)
//	fmt.Println(slice)  // cannot use arr (type [3]int) as type []int in argument to modify_slice
}