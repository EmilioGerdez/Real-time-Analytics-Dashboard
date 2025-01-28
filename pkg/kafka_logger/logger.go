package kafka_logger

import (
	"log"

	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/config"
	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/kafka"
)

type Logger struct {
	producer *kafka.Producer
}

func Init() (*Logger, error) {
	cfg := config.LoadKafkaConfig()
	producer, err := kafka.NewProducer(cfg.Brokers, cfg.Topic)
	if err != nil {
		return nil, err
	}

	// Redirect standard logger
	log.SetOutput(producer)
	log.SetFlags(0) // Remove standard flags if needed

	return &Logger{producer: producer}, nil
}

func (l *Logger) Close() error {
	return l.producer.Close()
}
