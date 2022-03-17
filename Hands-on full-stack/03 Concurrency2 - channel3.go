package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1) // receive from channel 1
		fmt.Println(<-c1) // receive from channel 2
	}

}

/* RESULT

Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms

*/
