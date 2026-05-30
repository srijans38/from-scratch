package helpers

import (
	"log"
	"net"

	"github.com/srijans38/from-scratch/http-server/internal/parsers"
)

func HandleConnection(c net.Conn) {
	defer c.Close()

	incomingData := make([]byte, 1024)
	c.Read(incomingData)

	log.Printf("Raw request data from %s: %s\n", c.RemoteAddr(), string(incomingData))

	request, err := parsers.ParseRequest(incomingData)

	if err != nil {
		log.Printf("Error parsing request from %s: %v\n", c.RemoteAddr(), err)

		c.Write([]byte("HTTP/1.1 400 Bad Request\r\n"))
		c.Write([]byte("Content-Type: text/plain\r\n\r\n"))
		c.Write([]byte("Error: " + err.Error()))

		return
	}

	log.Printf("Received %s request from %s\n", request.Method, c.RemoteAddr())

	c.Write([]byte("HTTP/1.1 200 OK\r\n"))
	c.Write([]byte("Content-Type: text/plain\r\n\r\n"))
	c.Write([]byte("Your request method was: " + string(request.Method)))
	c.Write([]byte("\nYour request path was: " + request.Path))
	c.Write([]byte("\nYour request had the following headers:\n"))

	for _, header := range request.Headers {
		c.Write([]byte(header.Key + ": " + header.Value + "\n"))
	}
}
