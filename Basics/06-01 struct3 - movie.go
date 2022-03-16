package main
import "fmt"

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
			name:   s,
			rating: r,
	}
	return
}

/*
func increaseRating(m *Movie) {
	m.rating += 1.0
}
   ---> cannot use mov (type Movie) as type *Movie in argument to increaseRating
*/

func increaseRating(m Movie) {
	m.rating += 1.0
}

func main() {
	mov := getMovie("xyz", 2.0)
	increaseRating(mov)
	fmt.Printf("%+v", mov)
}