package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func slow5() {
	// Create HTTP client
	client := &http.Client{}

	// Create HTTP request
	req, err := http.NewRequest("GET", "https://"+_url+"/path", nil)
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}

	// Set headers
	req.Header.Set("Host", _host)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Content-Length", "10000")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	//writer := bufio.NewWriter(body)

	// Keep connection alive
	for {
		// _, err := io.WriteString(writer, "X-a: b\r\n")
		// if err != nil {
		// 	// handle error
		// }

		//io.WriteString(resp.Body, "X-a: b\r\n")
		time.Sleep(time.Second)
		print("slow3")
	}

	// Keep connection alive
	for {
		// Create buffer with data to send
		//buf := bytes.NewBufferString("X-a: b\r\n")

		// Send data over connection
		// _, err := io.Copy(resp.Body, buf)
		// if err != nil {
		// 	// handle error
		// }

		time.Sleep(time.Second)
		print("slow3")
	}

	// Keep connection alive
	for {
		//fmt.Fprintf(resp.Body, "X-a: b\r\n")
		time.Sleep(time.Second)
		print("slow3")
	}

}
