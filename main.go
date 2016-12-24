package main

import (
	"bufio"
	"os"
	"log"
	"encoding/json"
	"flag"
	"net"
)

var (
	sendAddress = flag.String("sendAddress", "localhost:3001", "host:port the address to send to")
)

type Message struct {
	Body string
}

func main() {
	flag.Parse()

	connection, err := net.Dial("tcp", *sendAddress)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(connection)

	for reader.Scan() {
		message := Message{Body: reader.Text()}
		err := encoder.Encode(&message)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
