package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) (*LogMiddleware, error) {
	return &LogMiddleware{
		next: next,
	}, nil
}

func (l *LogMiddleware) ProduceData(data types.OBUData) {
	defer logrus.WithFields(
		logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
		},
	).Info("Producing to Kafka")

	l.next.ProduceData(data)
}
