package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// go worker(jobs, results) // <- 1 worker

	// MANY WORKERS - Taking advantage of the multi-core processing.
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// worker starts pulling concurrently a number at a time.
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs { // jobs: a queue of numbers, and the items from the queue will be consumed using the range feature.
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
