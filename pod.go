package main

import (
	"fmt"
	"net"
	"time"
)

const (
	maxICMPSize = 65535 // Maximum size of an ICMP packet
)

func main() {
	// Prompt the user for the target host and number of packets to send
	var host string
	var numPackets int
	fmt.Print("Enter target host: ")
	fmt.Scanln(&host)
	fmt.Print("Enter number of packets to send: ")
	fmt.Scanln(&numPackets)

	// Set up the ICMP connection
	conn, err := net.Dial("ip4:icmp", host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send the specified number of ICMP packets
	for i := 0; i < numPackets; i++ {
		fmt.Printf("Sending packet %d...\n", i+1)

		// Create an ICMP echo request packet
		p := make([]byte, maxICMPSize)
		p[0] = 8 // Echo request type
		p[1] = 0 // Code
		p[2] = 0 // Checksum (filled in later)
		p[3] = 0
		p[4] = 0 // Identifier
		p[5] = 13
		p[6] = 0 // Sequence number
		p[7] = 37

		// Fill the packet with random data
		for i := 8; i < maxICMPSize; i++ {
			p[i] = byte(rand.Intn(256))
	
