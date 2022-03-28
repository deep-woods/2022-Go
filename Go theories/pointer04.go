package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u // <-- returns the memory location.
}

func main() {
	userPointer := NewUser("Forest", 199)
	fmt.Printf("&userPointer  :%v \n", &userPointer)
	fmt.Printf("*userPointer  :%v \n", *userPointer)
	fmt.Printf("&userPointer  :%t \n", &userPointer)
	fmt.Printf("userPointer   :%v \n", userPointer)
}

/*

&userPointer  :0xc00000e028
*userPointer  :{Forest 199}
&userPointer  :%!t(**main.User=0xc00000e028)
userPointer   :&{Forest 199}

*/
