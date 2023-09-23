package main

import (
	"flare/internal/handler"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", "8080"))

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Print("Tcp server running on localhost:8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		connHanlder := handler.NewConnectionHandler()

		go connHanlder.Handle(conn)
	}
}
