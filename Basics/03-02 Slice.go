package main

import "fmt"

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[:4]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// slice and replace
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := arr1[:4]
	fmt.Println(arr1)
	fmt.Println(slice1)
	slice1[1] = "x"   // a simple replace operation
	fmt.Println(arr1)
	fmt.Println(slice1)

	/*
	- create an integer array of size 5 with values [10, 20, 90, 70, 60]
	- create a slice referencing the above array to contain elements from index 0, 1, 2
	- change the element at index 2 of slice to 900

	- print the array
	- print the slice */

	arr2 := [5]int{10, 20, 90, 70, 60}
	slice2 := arr2[:3]
	slice2[2] = 900

	fmt.Println(arr2)
	fmt.Println(slice2)
}

// https://gosamples.dev/capacity-and-length/