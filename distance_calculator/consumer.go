/*
	Purpose of this file:
		Consume incoming messages, calculate distance and send this distance to aggregator HTTP client.
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/sushant102004/Traffic-Toll-Microservice/dbAggregator/client"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type KafkaConsumer struct {
	consumer    *kafka.Conn
	calcService CalculatorServicer
	aggClient   *client.AggClient
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	calcService := NewCalculateService()
	aggClient := client.NewAggClient("http://localhost:3000/aggregate")

	return &KafkaConsumer{
		consumer:    conn,
		calcService: calcService,
		aggClient:   aggClient,
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

		dist = math.Floor(dist*100) / 100

		fmt.Println("Distance: ", dist)

		currentTime := time.Now()
		currentDate := currentTime.Format("02-01-2006")

		req := types.CalculatedDistance{
			OBUID:    data.OBUID,
			Distance: dist,
			Date:     currentDate,
		}

		if err := c.aggClient.PostDataToAPI(req); err != nil {
			logrus.Error("POST Error: ", err.Error())
		}
	}
}

func (c *KafkaConsumer) Start() {
	fmt.Println("Waiting For Data...")
	go c.consume()
}
