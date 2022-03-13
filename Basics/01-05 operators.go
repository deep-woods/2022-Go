package main
import (
	"fmt"
	"strconv"
)

func main(){
	var a, b string = "Monkeys", "monkeys"
	fmt.Println(a == b)

	var c string = "505"
	var d, err = strconv.Atoi(c)
	var e int = 505
	fmt.Println(e == d)
	fmt.Println(err)
	// ./operators.go:10:16: invali operation: c == d (mismatched types string and int)
}

