package main

import (
	"bufio"
	"os"
	"log"
	"encoding/json"
	"flag"
	"net"
	"fmt"
	"./util"
)

var (
	sendAddress = flag.String("sendAddress", ":3001", "host:port the address to send to")
	listenAddress string
)

type Message struct {
	Body string
	Address string
}

func main() {
	flag.Parse()

	messages := make(chan Message)
	go read(messages)
	go connect(messages)

	listen()
}

func listen() {
	listener, err := util.Listen()
	listenAddress = listener.Addr().String()
	log.Println("Listening on", listenAddress)
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

func read(messages chan Message) {
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		message := Message{Body: reader.Text(), Address: listenAddress}
		messages <- message
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}

func connect(messages chan Message) {
	connection, err := net.Dial("tcp", *sendAddress)
	if err != nil {
		log.Println(err)
		return
	}

	encoder := json.NewEncoder(connection)
	for message := range messages {
		err := encoder.Encode(&message)

		if err != nil {
			log.Println(err)
			return
		}
	}
}
