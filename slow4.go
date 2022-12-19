package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func slow4() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://thawdezin.netlify.app", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	// Send partial request
	//req.ContentLength = 10000

	// Keep connection alive
	for {
		req.Header.Set("X-a", "b")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		print("slow4")
		go slow3()
		time.Sleep(time.Second)
	}
}
