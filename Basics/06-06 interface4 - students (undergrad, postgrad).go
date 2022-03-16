package main

import "fmt"


type Student interface {
	getPercentage() int
	getName() string
}

type Undergrad struct {
	name   string
	grades []int
}

type Postgrad struct {
	name   string
	grades []int
}

func (p Postgrad) getPercentage() int {
	sum := 0
	for _, v := range p.grades {
			sum += v
	}
	return ((sum * 100) / (len(p.grades) * 200))
}
func (p Postgrad) getName() string {
	return p.name
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
			sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	u := Undergrad{"Ross", []int{90, 75, 80}}
	p := Postgrad{"Joe", []int{150, 190, 185}}
	printData(u)
	printData(p)
}

// Code from Kodekloud