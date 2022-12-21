package main

import (
	"fmt"
	"net"
)

//const _url = "www.example.com"
//const _host = "example.com"

func main2() {
	go slow3()
	select {}
}

func slow3_() {
	conn, err := net.Dial("tcp", _url+":443")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: "+_host+"\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	_, err = fmt.Fprintf(conn, "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	_, err = fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	_, err = fmt.Fprintf(conn, "Connection: keep-alive\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	_, err = fmt.Fprintf(conn, "Accept-Encoding: gzip, deflate\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}
	_, err = fmt.Fprintf(conn, "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n")
	if err != nil {
		fmt.Println(err)
		newErr := fmt.Errorf("Error Slow3: %v", err)
		fmt.Println(newErr)
		return
	}

	// Send partial request
	//_, err = fmt.
}