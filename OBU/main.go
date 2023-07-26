/*
	On Board Unit - This will send Location Coordinate to receiver service.

	Purpose of this code: -

	1. Generate random OBU IDs.
	2. Connect to WebSockets.
	3. Generate location coordinates based on OBU IDs.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const wsEndpoint = "ws://127.0.0.1:6443/ws"

func main() {

	obuID := rand.Intn(math.MaxInt)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		lat, long := generateLocation()
		data := types.OBUData{
			OBUID: obuID,
			Lat:   lat,
			Long:  long,
		}
		err := conn.WriteJSON(data)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("OBU ID: %d \nLat: %f \nLong: %f \n\n", data.OBUID, data.Lat, data.Long)
		time.Sleep(time.Second * 5)

	}
}

func generateCoord() float64 {
	n := float64(rand.Intn(100)) + 1
	f := rand.Float64()
	return n + f
}

func generateLocation() (float64, float64) {
	return generateCoord(), generateCoord()
}
