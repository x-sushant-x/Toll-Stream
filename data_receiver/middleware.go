/*
	Design Pattern: Decorator

	Purpose of this file -
	Apply decorator pattern to DataProducer interface and add a logger middleware.

	It will work something like next() function in express.js
*/

package main

import (
	"fmt"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

/*
next DataProducer refer to the original instance of DataProducer interface.
It will be used to call the original ProducerData function once our Logger had done it's work.
*/
type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) (*LogMiddleware, error) {
	return &LogMiddleware{
		next: next,
	}, nil
}

/*
	Write a function with name ProduceData which is defined in DataProducer interface.
	Write logging functionality and call LogMiddleware.next.ProduceData()
*/

func (l *LogMiddleware) ProduceData(data types.OBUData) {
	fmt.Println("Producing Data to Kafka")
	fmt.Println("OBU ID: ", data.OBUID)
	fmt.Println("Lat: ", data.Lat)
	fmt.Println("Long: ", data.Long)
	fmt.Println()

	l.next.ProduceData(data)
}
