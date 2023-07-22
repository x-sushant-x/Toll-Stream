/*
	Purpose of this file: -
	Start Consumer which is responsible for consuming and calculating distance from incoming kafka messages.
*/

package main

import "log"

func main() {
	kafkaConsumer, err := NewKafkaConsumer("toll-service")
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
	select {}
}
