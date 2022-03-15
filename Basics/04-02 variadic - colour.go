package main
import "fmt"

func nameColours(s string, colours ...string) {  // <--- notice the syntax of the variadic parameter. 
    fmt.Println(s)
    for _, value := range colours {
            fmt.Printf("%s, ", value)
    }
}

func main() {
    nameColours("Colours list", "ruby red", "diamond white", "skyblue", "city lights", "sunrise")
}
