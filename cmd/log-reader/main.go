package main

import (
	"context"
	"fmt"

	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/config"
	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/kafka"
	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/redis"
)

func main() {
	kcfg := config.LoadKafkaConfig()
	rcfg := config.LoadRedisConfig()

	// Connect to Redis
	rdb := redis.NewClient(rcfg)

	// Create Kafka consumer
	consumer, err := kafka.NewConsumer(kcfg.Brokers, kcfg.Topic)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Start consuming messages
	err = consumer.Read(context.Background(), func(message []byte) {
		// Store in Redis (example: last 100 logs)
		rdb.LPush(context.Background(), "recent_logs", string(message))
		rdb.LTrim(context.Background(), "recent_logs", 0, 99)

		fmt.Printf("Received log: %s\n", string(message))
	})
	if err != nil {
		panic(err)
	}
}
