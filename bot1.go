package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Bot struct {
	Client *http.Client
}

func NewBot() *Bot {
	return &Bot{
		Client: &http.Client{},
	}
}

func (b *Bot) PerformRequest(url string) error {
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

func main() {
	var wg sync.WaitGroup
	bots := []*Bot{}

	// Create a number of bots
	for i := 0; i < 10; i++ {
		bots = append(bots, NewBot())
	}

	// Control the bots to perform requests
	for _, bot := range bots {
		wg.Add(1)
		go func(bot *Bot) {
			defer wg.Done()
			if err := bot.PerformRequest("http://www.example.com/"); err != nil {
				fmt.Println(err)
			}

		}(bot)
	}

	wg.Wait()
}
