package main

import (
	"bufio"
	"os"
	"log"
	"encoding/json"
)

type Message struct {
	Body string
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

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
