package main

import "fmt"

type Student interface {
	getPercentage() int
	getName()
}

type Undergrad struct {
	name   string
	grades []int
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
		sum += v
	}
	return sum / len(u.grades)
}

func (u Undergrad) getName() {
	u.name = "Sylvia"
}

func printPercentage(s Student) {
	fmt.Println(s.getPercentage())
}

func main() {
	grades := []int{90, 75, 80}
	u := Undergrad{"Ross", grades}
	fmt.Printf("%+v\n", u)
	/* printPercentage(u) ---> cannot use u (type Undergrad) as type Student in argument to printPercentage:
        					   Undergrad does not implement Student (missing getName method)
	*/
	printPercentage(u)
	fmt.Printf("%+v\n", u)
} 