package main

import (
	"log"
	"net"

	"github.com/srijans38/from-scratch/http-server/internal/helpers"
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
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go helpers.HandleConnection(conn)
	}
}
