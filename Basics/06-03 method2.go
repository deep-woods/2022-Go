package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calcArea1() { // asterisk (*) modifies the original struct.
	c.area = 3.14 * c.radius * c.radius // {radius:45 area:6358.500000000001}
}

func (c Circle) calcArea2() { // does not modify the original struct value which is 0 by default. Hence prints 0 for area.
	c.area = 3.14 * c.radius * c.radius // {radius:45 area:0}
}

func main() {
	c1 := Circle{radius: 45}
	c1.calcArea1()
	fmt.Printf("%+v \n", c1)
	c2 := Circle{radius: 45}
	c2.calcArea2()
	fmt.Printf("%+v \n", c2)
}
