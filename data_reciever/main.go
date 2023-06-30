package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

func main() {
	dataReceiver := NewDataReceiver()
	http.HandleFunc("/ws", dataReceiver.wsHandler)

	if err := http.ListenAndServe(":30000", nil); err != nil {
		panic(err)
	}
}

type DataReceiver struct {
	msgCh chan types.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgCh: make(chan types.OBUData, 128),
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
	dr.conn = conn
	go dr.wsRecieveLoop()
}

func (dr *DataReceiver) wsRecieveLoop() {
	fmt.Println("Client Connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("Error: ", err)
			continue
		}
		dr.msgCh <- data
		res, ok := <-dr.msgCh
		if !ok {
			log.Println("Error while reading data from channel")
		}
		fmt.Println(res)
	}
}
