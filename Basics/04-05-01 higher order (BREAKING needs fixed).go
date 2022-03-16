package main
import ("fmt"
"math/rand")

/*
func createArray(n int) int[] {
	arr := []int{}
	for i := range n {
		arr = append(arr, rand.Intn(100))		
	}
	return arr
}
*/
func addHundred(x int) int {
	return x + 100
}
func partialSum(arr int[]) func() {
	sum := 0
	for _, value := range len(arr) {
			sum += value
	}
	return func() {
			fmt.Println(addHundred(sum))
	}
}
func main() {
	arr := []int{}
	for i := range 5 {
		arr = append(arr, rand.Intn(100))		
	}
	//arr := createArray(10)
	finalSum := partialSum(arr)
	finalSum()
}