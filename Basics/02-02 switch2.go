package main

import "fmt"

func main() {
        day := "sunday"
        switch day {
        case "monday":
                fmt.Println("monday")
        case "tuesday":
                fmt.Println("tuesday")
        case "wednesday":
                fmt.Println("wednesday")
        case "thursday":
                fmt.Println("thursday")
        case "friday":
                fmt.Println("friday")
        case "saturday", "sunday":
                fmt.Println("weekend")
                fallthrough
        default:
                fmt.Println("a whole week")
        }
        // "weekend"
        // "a whole week"
}
