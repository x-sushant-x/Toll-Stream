package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Conn
	isRunning   bool
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
		isRunning:   false,
		calcService: calcService,
	}, nil
}

func (c *KafkaConsumer) consume() {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:     []string{"localhost:9092"},
			Topic:       "toll-service",
			GroupID:     "toll-group",
			MinBytes:    1e3,
			MaxBytes:    10e3,
			StartOffset: kafka.LastOffset,
		},
	)

	defer reader.Close()

	go func() {
		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				logrus.Error("Error: ", err)
				continue
			}
			fmt.Println("Data: ", string(m.Value))
		}
	}()

	for {
		time.Sleep(time.Second * 1)
		hasNewMessage := reader.Stats().Messages > 0

		fmt.Println("Waiting for location coordinates")

		if hasNewMessage {
			break
		}
	}
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Started kafka consumer.")
	c.isRunning = true
	go c.consume()
}
