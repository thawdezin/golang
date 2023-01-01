package main

import (
	"fmt"
	"net/http"
)

func lbd1() {
	// Set up a client with the default transport
	client := http.Client{}

	// Make a request to a URL
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the "Connection" header to "close"
	req.Header.Set("Connection", "close")

	// Make the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Check the "Via" header in the response
	via := res.Header.Get("Via")
	if via != "" {
		fmt.Println("Load balancer detected!")
		fmt.Println("Via header:", via)
	} else {
		fmt.Println("No load balancer detected.")
	}
}
