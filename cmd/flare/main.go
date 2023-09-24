package main

import (
	"flare/internal/durable"
	"flare/internal/handler"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		durable.GetConfig(args[1])
	} else {
		durable.GetConfig("config.json")
	}

	config := durable.Config
	listener, err := net.Listen("tcp", net.JoinHostPort(config.Host, config.Port))

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Printf("Tcp server running on %v:%v", config.Host, config.Port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		connHanlder := handler.NewConnectionHandler()

		go connHanlder.Handle(conn)
	}
}
