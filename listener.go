package main

import (
	"net"
	"log"
	"fmt"
	"encoding/json"
	"flag"
)

var listenAddr = flag.String("listen", "localhost:3001", "host:port the address to listen on")

type Message struct {
	Body string
}

func serve(connection net.Conn) {
	defer connection.Close()
	decoder := json.NewDecoder(connection)
	for {
		var message Message
		err := decoder.Decode(&message)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", message)
	}
}

func main() {
	listener, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go serve(connection)
	}
}