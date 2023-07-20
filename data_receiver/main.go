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
	prod    DataProducer
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

func (dr *DataReceiver) wsRecieveLoop() {
	fmt.Println("Client Connected")
	for {
		var data types.OBUData
		if err := dr.webSCon.ReadJSON(&data); err != nil {
			log.Println("Error: ", err)
			continue
		}
		dr.prod.ProduceData(data)
	}
}
