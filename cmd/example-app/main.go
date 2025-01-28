package main

import (
	"log"

	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger"
)

func main() {
	logger, err := kafka_logger.Init()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	log.Println("This log will be sent to Kafka")
}
