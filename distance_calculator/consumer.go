package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Conn
	isRunning bool
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return &KafkaConsumer{
		consumer: conn,
	}, nil
}

func (c *KafkaConsumer) consume() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(1e3)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(msg.Value))

		b := make([]byte, 10e3)
		_ = b
	}
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Started kafka consumer.")
	c.isRunning = true
	go c.consume()
}
