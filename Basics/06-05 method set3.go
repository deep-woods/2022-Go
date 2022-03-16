package main

import "fmt"

type Employee struct {
	dept string
	eid  int
}

func (e *Employee) get_id(n int) {
	e.eid = n
}

func (e *Employee) get_dept(dept string) {
	e.dept = dept
}

func main() {
	employees := make([]Employee, 5)
	for i := range employees {
		fmt.Printf("%+v ", employees[i])
		employees[i].get_id(i)
		employees[i].get_dept("dev")
		fmt.Printf("%+v\n", employees[i])
	}
}
