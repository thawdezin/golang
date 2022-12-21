package main

import "runtime"

func recursiveFunction() {
    // Get information about the system's memory usage
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)

	go slow4()
    // Check the system's available memory
    if mem.Alloc > 500000000 { // 1 GB
        // Make the next recursive call in a new goroutine
        go recursiveFunction()
    } else {
        // Stop making recursive calls
		print("========== RAM FULL ==========")
        return
    }
}
