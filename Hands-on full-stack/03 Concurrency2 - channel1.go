package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	go count("gazelle", c)

	// VERSION 2
	for msg := range c { // <------ counting the range of the channel so close check is not necessary.
		fmt.Println(msg)
	}

	/*	VERSION 1
		for {
			msg, open := <-c

			if !open {        <<------- Manually checking if the channel is closed or not.
				break
			}
			fmt.Println(msg)
		}
	*/
}

func count(thing string, c chan string) {

	for i := 1; i <= 7; i++ {
		c <- thing + " " + strconv.Itoa(i)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
