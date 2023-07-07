package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const topic = "toll-calculator"
const partition = 0

type DataProducer interface {
	ProduceData(types.OBUData) error
}

type KafkaProducer struct {
	producer *kafka.Conn
}

func NewKafkaProducer() *KafkaProducer {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return &KafkaProducer{
		producer: conn,
	}
}

func (dr KafkaProducer) ProduceData(data types.OBUData) error {
	d, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	_, err = dr.producer.WriteMessages(
		kafka.Message{Value: []byte(d)},
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return nil
}
