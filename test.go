package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runFunction()
		}()
	}

	wg.Wait()
}

func runFunction() {
	for {
		// code to run function goes here

		slow4()
		time.Sleep(1 * time.Second)
	}
}