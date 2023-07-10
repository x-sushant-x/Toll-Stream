package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const topic = "toll-calculator"
const partition = 0

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	batch := conn.ReadBatch(1, 1e6)

	b := make([]byte, 10e3)

	for {
		n, err := batch.Read(b)
		if err != nil {
			fmt.Println("Error reading batch data.")
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch.")
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection.")
	}
}
