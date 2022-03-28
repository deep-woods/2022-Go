package custompkg

// !(package main) == this file is not the entrypoint, not a programme.
// 				   == this is an auxiliary package for the main programme.

// C:/User/go/pkg/ ... you can find all the librariesi you've downloaded here.
// Command "go evn" for short help on go environment conifguration.

import "fmt"

func PrintCustom() {
	fmt.Println("This is a custom package!")
}
