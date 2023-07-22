/*
	Purpose of this file -
	1. Receive OBU Data and produce this data to Kafka.

	Important: - Producing data means adding incoming data to Kafka Message Queue
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

func main() {
	dataReceiver, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	// Providing function to execute on PORT 6643 to HTTP.
	http.HandleFunc("/ws", dataReceiver.wsHandler)

	go func() {
		if err := http.ListenAndServe(":6443", nil); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Waiting for Coordinates...")
	select {}
}

type DataReceiver struct {
	webSCon *websocket.Conn
	// DataProducer interface contains funciton to produce data to Kafka.
	prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		prod DataProducer
		err  error
	)

	prod, err = NewKafkaProducer()
	if err != nil {
		return nil, err
	}

	prod, err = NewLogMiddleware(prod)
	if err != nil {
		return nil, err
	}

	return &DataReceiver{
		prod: prod,
	}, nil
}

// This function will handle all the incoming WebSocket messages.
func (dr *DataReceiver) wsHandler(w http.ResponseWriter, req *http.Request) {
	/*
		Upgrading HTTP Connection to Websocket Connection.
	*/
	u := websocket.Upgrader{ReadBufferSize: 1028, WriteBufferSize: 1028}

	conn, err := u.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.webSCon = conn

	go dr.wsRecieveLoop()
}

// Continously read received messages from the websocket connection.
func (dr *DataReceiver) wsRecieveLoop() {
	fmt.Println("Client Connected")
	for {
		var data types.OBUData
		if err := dr.webSCon.ReadJSON(&data); err != nil {
			log.Println("Error: ", err)
			continue
		}

		// Producing to Kafka using DataReceiver struct's
		dr.prod.ProduceData(data)
	}
}
