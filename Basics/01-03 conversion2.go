package main
import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "505I'mgoingbackto505"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)
}

// Result
// 0, int 
// strconv.Atoi: parsing "505I'mgoingbackto505": invalid syntax, *strconv.NumError 