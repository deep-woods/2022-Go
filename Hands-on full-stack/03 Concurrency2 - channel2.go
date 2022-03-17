package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "Hello" // send will block until something is ready to receive.
	c <- "world"
	// c <- "three" ------------------ When you run this,

	msg := <-c // receive lines
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

/*  ------------------ You get this.
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
		/home/ai-square/WORKSPACE/Hands-on full-stack/03 Concurrency2 - channel2.go:11 +0x8d
exit status 2
*/
