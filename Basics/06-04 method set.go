package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	s := Student{name: "Alex", grades: []int{35, 86, 99}}
	s.displayName()
	fmt.Printf("%.2f%% \n", s.calculatePercentage())
}

// code from Kodekloud
