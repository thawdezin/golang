package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
if err != nil {
	fmt.Printf("Error creating request: %s\n", err)
}
req.Header.Add("User-Agent", "My Go Client 1.0")

response, err := http.DefaultClient.Do(req)
if err != nil {
	fmt.Printf("The HTTP request failed with error %s\n", err)
} else {
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

}
