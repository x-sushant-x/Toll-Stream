/*
	Purpose of this file -
	Define DataProducer interface and a ProduceData function inside it.

	ProduceData function will add received messages to Kafka Queue.

	Dependencies - Kafka Connection
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const topic = "toll-service"
const partition = 0

/*
We are using interface instead of directly using struct because this
will be passed through decorator pattern middleware which will first
print out logs and than call this ProduceData method.

That middleware will be of LoggerMiddleware types and our main Kafka
Producer will be of KafkaProducer type.

To prevent mismatch of types we can use DataProducer interface.
*/
type DataProducer interface {
	ProduceData(types.OBUData)
}

type KafkaProducer struct {
	producer *kafka.Conn
}

func NewKafkaProducer() (*KafkaProducer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return &KafkaProducer{
		producer: conn,
	}, nil
}

func (dr KafkaProducer) ProduceData(data types.OBUData) {
	d, err := json.Marshal(&data)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	_, err = dr.producer.WriteMessages(
		kafka.Message{Value: []byte(d)},
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
