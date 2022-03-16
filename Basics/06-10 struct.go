package main

import "fmt"

type band struct {
	name string
	songtitle string
	members []string
}

// Initialisation method 1 
var arctic_monkeys = band {
	name: "Architic Monkeys",
	songtitle: "505",
	members: []string{"Alex", "Matt", "Jamie", "Nick"},
}

// Initialisation method 2 - struct literal 
var capital_cities = band{"Capital Cities", "Kangaroo Court", []int{"Sebu", "Ryan"}}

// Initialisation method 3 - dot notation
var elodie = band{}
elodie.name = "Elodie"
elodie.songtitle = "Margarita"
elodie.members = []string{"Elodie"}

func main() {
	fmt.Println(arctic_monkeys)
	fmt.Println(capital_cities)
	fmt.Println(elodie)
}