package main

import "fmt"

func main() {
        var i, j = 7, 45

        switch {  		// << no switch target is specified. 
        case i+j < 60:  // this expression is possible because the switch target is not specified in the preceding line. 
                fmt.Println("less than 60")
        case i+j = 60:
                fmt.Println("equal to 60")
                fallthrough
        default:
                fmt.Println("greater than 60")
        }

}