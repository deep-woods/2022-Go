package main

import "fmt"

func main() {
	// Map array element to a map.
	arr := [5]string{"one", "two", "three"}
	slice := arr[:3]
	my_map := make(map[int]string)
	for i, el := range slice {  // map it using for loop.
			my_map[i+1] = el   
	}
	fmt.Println(my_map)

	// func make to create maps and arrays. 
	arr := [10]string{"a", "b", "c"}
	hashmap := make(map[string]int)
	my_slice := arr[:5]
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))
	fmt.Println(len(hashmap))
}