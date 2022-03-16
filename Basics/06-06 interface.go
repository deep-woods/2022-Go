package main

import "fmt"

/* type <interface name> interface {

	1. A type implements an interface by implementing its methods. 
	2. Go interfaces are implemented implicitly.
	3. Go interfaces do not have any specific keyword to implement them.

} */

type FixedDeposite interface {
	getRateOfInterest() float64
	calcReturn() float64
}