package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// strings.ToUpper
	x := func(s string) string {
			return strings.ToUpper(s)
	}
	fmt.Printf("%T \n", x)	// func(string) string 
	fmt.Println(x("Cloe"))	// CLOE

	// strconv.Itoa
	var (
		cube = func(i int) string {
				c := i * i * i
				return strconv.Itoa(c)
		}
	)
	
	x := cube(8)
	fmt.Printf("%T %v", x, x) // string 512
}
