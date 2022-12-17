package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// Make an HTTP GET request to the website
	response, err := http.Get("http://thawdezin.netlify.app")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
