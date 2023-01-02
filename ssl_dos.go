package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	// Prompt the user for the target host and number of connections to establish
	var host string
	var numConnections int
	fmt.Print("Enter target host: ")
	fmt.Scanln(&host)
	fmt.Print("Enter number of connections to establish: ")
	fmt.Scanln(&numConnections)

	// Establish the specified number of SSL connections
	for i := 0; i < numConnections; i++ {
		fmt.Printf("Establishing connection %d...\n", i+1)

		// Set up the TLS connection
		conn, err := tls.Dial("tcp", host, &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Close the connection immediately after establishing it
		conn.Close()
	}
}
