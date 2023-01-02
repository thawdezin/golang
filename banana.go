package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Banana struct {
	Client *http.Client
}

func NewBanana() *Banana {
	return &Banana{
		Client: &http.Client{},
	}
}

func (b *Banana) PerformRequest(url string) error {
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return err
	}

	resp, err := b.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (b *Banana) Run() {
	// Run the Banana app in the background and listen for orders from the Orange app
	for {
		// Wait for an order from the Orange app
		order := WaitForOrder()

		// Perform the requested action
		switch order {
		case "attack":
			// Send an HTTP request to a specified website
			url := GetTargetURL()
			if err := b.PerformRequest(url); err != nil {
				fmt.Println(err)
			}
		case "say":
			// Print "Hi"
			fmt.Println("Hi")
		}
	}
}

func main() {
	// Create a new Banana app and run it
	banana := NewBanana()
	banana.Run()
}
