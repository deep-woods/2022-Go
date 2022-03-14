package main

import "fmt"

func main() {
	arr := [5]bool{true, true, true, true}
	for i := 0; i < len(arr); i++ {
		if arr[i] {
				fmt.Println(i) // if the element is true, print index. 
		}
	}

	arr2 := [5]string{"a", "b", "c"}
	for index, element := range arr2 {
			fmt.Println(index, "->", element)
	}

	arr3 := [5][2]string{{"a"}, {"b"}, {"c"}} // [[a ] [b ] [c ] [ ] [ ]]
	fmt.Println(arr3[0][0])
	fmt.Println(arr3[1][1])
	fmt.Println(arr3[2][0])
}