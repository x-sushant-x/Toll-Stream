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
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(1e3)
		if err != nil {
			fmt.Print(err)
		}
		// fmt.Println(string(msg.Value))

		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Error("Error: ", err)
			continue
		}
		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Error("Error: Unable to parse data.")
			continue
		}
		fmt.Println("Distance: ", distance, " KMs")
	}
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Started kafka consumer.")
	c.isRunning = true
	go c.consume()
}
