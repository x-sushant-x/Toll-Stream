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
	// coords := []float32{
	// 	29.39364376723343, 76.96343000761095, 29.298344310716402, 76.99645257497889, 29.241565027084423, 77.0108110740307, 29.1425395225112, 77.03918699127212, 28.92469795054101, 77.10214459503118, 28.688381264122476, 77.2160040910087,
	// }

	// obuIDS := generateOBUIDs(20)

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

	// for {
	// 	for i := 0; i < len(obuIDS); i++ {
	// 		lat, long := generateLocation()
	// 		data := types.OBUData{
	// 			OBUID: obuIDS[i],
	// 			Lat:   lat,
	// 			Long:  long,
	// 		}
	// 		err := conn.WriteJSON(data)
	// 		if err != nil {
	// 			fmt.Println("Error", err)
	// 		}

	// 		fmt.Printf("OBU ID: %d \nLat: %f \nLong: %f \n\n", data.OBUID, data.Lat, data.Long)
	// 		time.Sleep(time.Second * 5)
	// 	}
	// }
}

// func init() {
// 	/* If this is not done then program will generate same set of random numbers. */
// 	rand.New(rand.NewSource(time.Now().UnixMilli()))
// }

func generateCoord() float64 {
	n := float64(rand.Intn(100)) + 1
	f := rand.Float64()
	return n + f
}

func generateLocation() (float64, float64) {
	return generateCoord(), generateCoord()
}
