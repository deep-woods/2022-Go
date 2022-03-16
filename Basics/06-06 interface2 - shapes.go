package main

import "fmt"

type square struct {
	side float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}


type rect struct {
	length, breath float64
}

func (r rect) area() float64 {
	return r.length * r.breath
}

func (r rect) perimeter() float64 {
	return 2 * r.length + 2 * r.breath
}
 
func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	s := square{4}
	r := rect{7, 3}
	printData(s)
	printData(r)
}