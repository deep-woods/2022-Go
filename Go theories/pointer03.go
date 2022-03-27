package main

import (
	"fmt"
)

type Data struct {
	value int
	arr   [45]int
}

func ChangeData1(arg Data) {
	arg.value = 505
	arg.arr[7] = arg.value

	fmt.Println("Inside the ChangeData1 func")
	fmt.Printf("value = %d\n", arg.value)
	fmt.Printf("arr[7] = %d\n\n", arg.arr[7])
}

func ChangeData2(arg *Data) {
	arg.value = 505
	arg.arr[7] = arg.value

	fmt.Println("Inside the ChangeData2 func")
	fmt.Printf("value = %d\n", arg.value)
	fmt.Printf("arr[7] = %d\n\n", arg.arr[7])
}

func main() {
	var (
		data1 Data
		data2 Data
	)

	ChangeData1(data1) // <--- data1 and arg in ChanData(arg) are two disparate memory spaces.
	fmt.Println("Outside the ChangeData1 func")
	fmt.Printf("value = %d\n", data1.value)
	fmt.Printf("arr[7] = %d\n\n", data1.arr[7])

	ChangeData2(&data2)
	fmt.Println("Outside the ChangeData2 func")
	fmt.Printf("&data2: %v\n", &data2)
	fmt.Printf("*data2: %s\n", "invalid operation")
	fmt.Printf("value = %d\n", data2.value)
	fmt.Printf("arr[7] = %d\n", data2.arr[7])
}

/*

Inside the ChangeData1 func
value = 505
arr[7] = 505

Outside the ChangeData1 func
value = 0
arr[7] = 0

Inside the ChangeData2 func
value = 505
arr[7] = 505

Outside the ChangeData2 func
&data2: &{505 [0 0 0 0 0 0 0 505 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
*data2: invalid operation
value = 505
arr[7] = 505

*/
