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
