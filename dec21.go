package main

import (
    "net/http"
    "time"
)

func dec21() {
    // Create a new HTTP client
    client := &http.Client{}

    // Create a new HTTP request
    req, err := http.NewRequest("GET", _url, nil)
    if err != nil {
        print(err)
    }

    // Set the Host and User-Agent headers
    req.Header.Set("Host", _url)
    req.Header.Set("User-Agent", "MyGoClient/1.0")

    // Send the request with the first two headers
    _, err = client.Do(req)
    if err != nil {
        print("err1")
    }

	


    // Wait for 3 seconds
    time.Sleep(3 * time.Second)

    // Add additional headers to the request
    req.Header.Set("X-Custom-Header", "value")
    req.Header.Set("X-Another-Header", "value")

    // Send the request with the additional headers
    _, err = client.Do(req)
    if err != nil {
        print("err2")
    }

    // Wait for 5 seconds
    time.Sleep(5 * time.Second)

    // Process the response
    // (omitted for brevity)
}
