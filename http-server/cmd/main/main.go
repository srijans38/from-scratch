package main

import (
	"log"
	"net"
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
			log.Printf("Received data from %s: %s\n", c.RemoteAddr(), string(incomingData))
			c.Write([]byte("Hello, World!\n"))
			log.Printf("Sent greeting to %s\n", c.RemoteAddr())
		}(conn)
	}
}
