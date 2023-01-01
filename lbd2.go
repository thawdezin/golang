package main

import (
	"fmt"
	"net/http"
)

func checkLoadBalancer(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	if realIP := resp.Header.Get("X-Real-IP"); realIP != "" {
		return true
	}
	return false
}

func main() {
	fmt.Println("Enter a URL:")
	var url string
	fmt.Scanln(&url)

	if checkLoadBalancer(url) {
		fmt.Println("The given URL is behind a load balancer.")
	} else {
		fmt.Println("The given URL is not behind a load balancer.")
	}
}
