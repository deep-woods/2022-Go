package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) incLength(n int) {
	for i := 0; i < n; i++ {
		r.length += i
	}
}

func main() {
	r := Rectangle{breadth: 10, length: 5}
	fmt.Println(r.area()) // 50
	fmt.Println(r)        // {5 10}
	r.incLength(7)
	fmt.Println(r.area()) // 260
	fmt.Println(r)        // {26 10}
}
