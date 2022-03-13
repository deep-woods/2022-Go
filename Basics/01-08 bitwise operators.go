package main
import "fmt"

// bitwise AND (&)
// bitwise OR (|)
// bitwise XOR (^)
// left shift (<<)
// right shift (>>)

func main() {
	var x, y int = 100,90
	fmt.Println(x & y)  // 64 (1000000)
	fmt.Println(x | y)  // 126 (1111110) = 64 + 32 + 16 + 8 + 4 + 2
	// x = 1100100 
	// y = 1011010  

	var a, b int = 12, 25
	c := a & b
	fmt.Println(c)
	// a =  1100
	// b = 11001
	// c =  1000 = 8(10)

	var e, f int = 13, 57
	g := e ^ f
	fmt.Println(g)
	// e =   1101
	// f = 111001
	// g = 110100 = 52(10)

	var j int = 212
	k := j >> 2     // j = 11010100
	fmt.Println(k)  // k = 110101/00 = 53(10) 

	var m, n int = 100,90
	fmt.Println(!(((m + n) >> 2 ) == 47))
	// 190 = 128 32 16 8 4 2 = 10111110
	// 190 >> 2 = 101111/10 = 47
}

