package main

import "fmt"

type Employee struct {
	name   string
	salary int
	dept   string
	level  int
}

func (e Employee) getDept() { // doesn't change the original value.
	e.dept = "dev"
}

func (e *Employee) increaseSalary(bonus int) { // askterisk (*) modifies the original value.
	e.salary += bonus
}

func main() {
	emp := Employee{name: "Florence", salary: 3000}
	fmt.Printf("%+v \n", emp) // {name:Florence salary:3000 dept: level:0}
	emp.increaseSalary(200)
	emp.getDept()
	fmt.Printf("%+v", emp) // {name:Florence salary:3200 dept: level:0}
}
