package main
import "fmt"

func modify(m map[string]int) {
	m["place"] = 505
}

func main() {
	ascii_codes := make(map[string]int)
	ascii_codes["drive"] = 45
	ascii_codes["flight"] = 7
	fmt.Println(ascii_codes)
	modify(ascii_codes)
	fmt.Println(ascii_codes)
}