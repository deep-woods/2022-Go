package main
import "fmt"

func returns_a_printing_func() func(string) {
    return func(message string){
        fmt.Println(message)
    }
}

func main() {
	func_var := returns_a_printing_func()
	func_var("I'm going back to 505.")
}