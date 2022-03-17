package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.WaitGroup()

	/*
		go count("cheetah")
		go count("lion")
		// time.Sleep(time.Second * 5)
		fmt.Scanln() // blocking call to wait for user input.
	*/
}
