package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type KafkaConsumer struct {
	consumer    *kafka.Conn
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	calcService := NewCalculateService()

	return &KafkaConsumer{
		consumer:    conn,
		calcService: calcService,
	}, nil
}

func (c *KafkaConsumer) consume() {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "toll-service",
			GroupID:  "toll-group",
			MinBytes: 1e3,
			MaxBytes: 10e3,
		},
	)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error("Error: ", err)
			continue
		}
		var data types.OBUData
		if err = json.Unmarshal(m.Value, &data); err != nil {
			logrus.Error("Serialization Error: ", err)
		}

		dist, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Error("Calculate Error: ", err)
		}
		fmt.Println("Distance: ", dist)
	}
}

func (c *KafkaConsumer) Start() {
	fmt.Println("Waiting For Data...")
	go c.consume()
}
