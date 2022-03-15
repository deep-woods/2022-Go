package main
import "fmt"

func printStrings(names ...string) (names_c []string) {
    names_c := []string {}   // no new variables on left side of :=
        for _, value := range names {
                names_c = append(names_c, strings.ToUpper(value))
        }                 // string.ToUpper undefined (type string has no method ToUpper)
        return
}

func main() {
        result := printStrings("Joe", "Monica", "Gunther")
        fmt.Println(result)
}
