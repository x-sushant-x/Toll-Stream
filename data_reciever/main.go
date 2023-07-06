package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const topic = "toll-calculator"
const partition = 0

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	dataReceiver := NewDataReceiver(conn)
	http.HandleFunc("/ws", dataReceiver.wsHandler)

	if err := http.ListenAndServe(":30000", nil); err != nil {
		panic(err)
	}
}

type DataReceiver struct {
	webSCon *websocket.Conn
	conn    *kafka.Conn
}

func NewDataReceiver(conn *kafka.Conn) *DataReceiver {
	return &DataReceiver{
		conn: conn,
	}
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

		d, err := json.Marshal(&data)
		if err != nil {
			log.Fatal(err)
		}

		_, err = dr.conn.WriteMessages(
			kafka.Message{Value: []byte(d)},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
