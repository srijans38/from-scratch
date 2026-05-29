package main

import (
	"log"
	"net"

	"github.com/srijans38/from-scratch/http-server/cmd/internal/parsers"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("Listening on :8080")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(c net.Conn) {
			incomingData := make([]byte, 1024)
			defer c.Close()
			c.Read(incomingData)
			method, err := parsers.ParseRequest(incomingData)
			if err != nil {
				log.Printf("Error parsing request from %s: %v\n", c.RemoteAddr(), err)
				c.Write([]byte("HTTP/1.1 400 Bad Request\r\n"))
				c.Write([]byte("Content-Type: text/plain\r\n\r\n"))
				c.Write([]byte("Error: " + err.Error()))
				return
			}
			log.Printf("Received %s request from %s\n", method, c.RemoteAddr())
			c.Write([]byte("HTTP/1.1 200 OK\r\n"))
			c.Write([]byte("Content-Type: text/plain\r\n\r\n"))
			c.Write([]byte("Your request method was: " + method))
		}(conn)
	}
}
