package main

import (
	"runtime"
	"time"
)

func main() {
	// Get the maximum number of CPUs that can be executing simultaneously
	maxProcs := runtime.GOMAXPROCS(0)

	// Start a goroutine for each CPU to print "Hello"
	for i := 0; i < maxProcs; i++ {
		go func() {
			for {
				slow2()
			}
		}()
	}

	// Sleep indefinitely to keep the program running
	for {
		time.Sleep(time.Hour)
	}
}
