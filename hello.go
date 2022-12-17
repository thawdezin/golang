package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// req, err := http.NewRequest("GET", "https://thawdezin.netlify.app", nil)
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

	// req, err := http.NewRequest("GET", "https://thawdezin.netlify.app", nil)
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
	// response, err := client.Get("https://thawdezin.netlify.app")
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }

	// Set up a client with a long timeout
	client := &http.Client{
		Timeout: time.Hour,
	}

	// Send a partial HTTP request to the server and keep the connection open
	for {
		req, err := http.NewRequest("GET", "https://thawdezin.netlify.app", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("User-Agent", "!")

		// Send the request and ignore the response
		_, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			//data, _ := ioutil.ReadAll(response.Body)
			fmt.Println("result OK")
		}

		// Sleep for a short time before sending the next partial request
		time.Sleep(time.Second)
	}

}
